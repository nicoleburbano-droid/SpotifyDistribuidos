package main

import (
	controlador "almacenamiento/CapaControladores"
	capafachada "almacenamiento/CapaFachadaServices/Fachada"
	"fmt"
	"net/http"
	"os"
)

// Servidor de audios: el administrador sube por REST nuevos mp3. El archivo
// se guarda donde lo lee el servidor de streaming y los metadatos se
// registran en el servidor de metadatos.
func main() {
	puerto := obtenerVariable("PUERTO_AUDIOS", "5000")
	// El servidor de streaming abre los audios desde su propia carpeta.
	carpetaAudios := obtenerVariable("CARPETA_AUDIOS", "../servidorStreaming")
	urlMetadatos := obtenerVariable("URL_METADATOS", "http://localhost:8080")

	fachada := capafachada.NuevaFachadaAlmacenamiento(carpetaAudios, urlMetadatos)
	ctrl := controlador.NuevoControladorAlmacenamientoAudios(fachada)

	http.HandleFunc("/audios/almacenamiento", ctrl.AlmacenarAudio)

	fmt.Printf("Servidor de audios escuchando en el puerto %s (guarda en %s)...\n", puerto, carpetaAudios)
	if err := http.ListenAndServe(":"+puerto, nil); err != nil {
		fmt.Println("Error iniciando el servidor:", err)
	}
}

// obtenerVariable lee una variable de entorno o retorna el valor por defecto.
func obtenerVariable(nombre string, valorPorDefecto string) string {
	if valor := os.Getenv(nombre); valor != "" {
		return valor
	}
	return valorPorDefecto
}
