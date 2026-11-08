package main

import (
	repositorio "almacenamiento/CapaAccesoADatos/repository"
	controlador "almacenamiento/CapaControladores"
	clienteMetadatos "almacenamiento/CapaFachadaServices/ComponenteClienteMetadatos"
	capafachada "almacenamiento/CapaFachadaServices/Fachada"
	"almacenamiento/discovery"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const nombreServicio = "servidorDeAudios"

// Servidor de audios: el administrador sube por REST nuevos mp3. El archivo
// se guarda donde lo lee el servidor de streaming y los metadatos se
// registran en el servidor de metadatos.
func main() {
	puerto := obtenerVariable("PUERTO_AUDIOS", "5000")
	numeroPuerto, err := strconv.Atoi(puerto)
	if err != nil || numeroPuerto < 1 || numeroPuerto > 65535 {
		log.Fatalf("PUERTO_AUDIOS inválido: %q", puerto)
	}
	// El servidor de streaming abre los audios desde su propia carpeta.
	carpetaAudios := obtenerVariable("CARPETA_AUDIOS", "../servidorStreaming")
	metadataAudioRepository := repositorio.NewMetadataAudioRepository()
	serviceMetadatos := clienteMetadatos.NewMetadataAudioService(metadataAudioRepository)

	fachada := capafachada.NuevaFachadaAlmacenamiento(carpetaAudios, serviceMetadatos)
	ctrl := controlador.NuevoControladorAlmacenamientoAudios(fachada)

	router := gin.Default()
	metadataController := controlador.NewMetadataAudioController(serviceMetadatos)

	// Almacenamiento de archivos recibido desde el administrador.
	router.POST("/audios/almacenamiento", gin.WrapF(ctrl.AlmacenarAudio))

	// Consultas de metadata servidas por el segundo controlador.
	router.GET("/audios/titulo/:titulo", metadataController.ConsultarAudio)
	router.GET("/audios/tipo/:tipo", metadataController.ConsultarAudiosPorTipo)
	router.GET("/musica/titulo/:titulo", metadataController.ConsultarMusica)
	router.GET("/musica/all", metadataController.ObtenerMusica)
	router.GET("/audiolibros/titulo/:titulo", metadataController.ConsultarAudioLibro)
	router.GET("/audiolibros/all", metadataController.ObtenerAudioLibros)
	router.GET("/podcasts/titulo/:titulo", metadataController.ConsultarPodcast)
	router.GET("/podcasts/all", metadataController.ObtenerPodcast)
	router.GET("/ruido-blanco/tipo/:tipoSonido", metadataController.ConsultarRuidoBlanco)
	router.GET("/ruido-blanco/all", metadataController.ObtenerRuidoBlanco)
	router.GET("/tipos-audio", metadataController.ObtenerTiposDeAudio)

	servidor := &http.Server{Addr: ":" + puerto, Handler: router}
	go func() {
		fmt.Printf("Servidor de audios escuchando en el puerto %s (guarda en %s)...\n", puerto, carpetaAudios)
		if err := servidor.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error iniciando el servidor: %v", err)
		}
	}()

	registro, err := discovery.Registrar(obtenerVariable("CONSUL_ADDR", "127.0.0.1:8500"), nombreServicio, obtenerVariable("SERVICE_HOST", "127.0.0.1"), numeroPuerto)
	if err != nil {
		log.Fatalf("no fue posible registrarse en Consul: %v", err)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	registro.Desregistrar()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = servidor.Shutdown(ctx)
}

// obtenerVariable lee una variable de entorno o retorna el valor por defecto.
func obtenerVariable(nombre string, valorPorDefecto string) string {
	if valor := os.Getenv(nombre); valor != "" {
		return valor
	}
	return valorPorDefecto
}
