package vistas

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"cliente.local/grpc-cliente/dtos"

	util "cliente.local/grpc-cliente/utilidades"
	pb "servidorStreaming.local/grpc-servidor/serviciosAudio"
)

/*
Funcion que lee el titulo de un audio que se quiere reproducir
e invoca un procedimiento remoto que se contecta al servidor de audios
*/
func RecibirAudioAReproducir(client pb.AudioServiceClient, ctx context.Context) {

	readerInput := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el título del audio: ")
	titulo, _ := readerInput.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	//Invocación del procedimiento remoto
	stream, err := client.AudioStream(ctx, &pb.AudioRequest{Filename: titulo})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Recibiendo y reproduciendo audio en vivo...")
	reader, writer := io.Pipe()
	canalSincronizacion := make(chan struct{})

	// Arranca la goroutine de decodificación y reproducción de los fragmentos
	go util.DecodificarReproducir(reader, canalSincronizacion)

	// Arranca la recepción de los fragmentos desde el servidor
	util.RecibirAudio(stream, writer, canalSincronizacion)

}

/*
Fachada REST para pedirle al servidor de metadatos los audios
disponibles para un determinado tipo y mostrar sus títulos.
*/
func MostrarAudiosPorTipo(tipo string, ctx context.Context) {

	// Escapar el tipo para poder utilizarlo correctamente
	// como parámetro de la URL.
	tipoEscapado := url.PathEscape(tipo)

	serverURL := fmt.Sprintf(
		"http://localhost:8080/audios/%s",
		tipoEscapado,
	)

	// Crear cliente HTTP.
	httpClient := &http.Client{}

	// Crear la petición GET asociada al contexto recibido.
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		serverURL,
		nil,
	)
	if err != nil {
		fmt.Printf("Error al crear la petición HTTP: %v\n", err)
		return
	}

	// Realizar la petición REST.
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error al conectar con el servidor de metadatos: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// El servidor indica que no se encontraron audios.
	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Printf(
			"No se encontraron audios para el tipo '%s'.\n",
			tipo,
		)
		return
	}

	// Cualquier código diferente de 200 se considera
	// una respuesta no exitosa.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf(
			"Error al consultar los audios. Código HTTP: %d\n",
			resp.StatusCode,
		)
		return
	}

	// DTO que representa exactamente la respuesta del servidor REST.
	var respuesta dtos.RespuestaAudiosPorTipoDTO

	// Convertir el JSON recibido a una estructura Go.
	err = json.NewDecoder(resp.Body).Decode(&respuesta)
	if err != nil {
		fmt.Printf(
			"Error al interpretar la respuesta JSON: %v\n",
			err,
		)
		return
	}

	fmt.Printf("\n%s\n", respuesta.Mensaje)

	// Verificar si existen audios.
	if len(respuesta.VectorAudiosPorTipo) == 0 {
		fmt.Printf(
			"No hay audios disponibles para el tipo '%s'.\n",
			tipo,
		)
		return
	}

	// Mostrar únicamente los títulos.
	fmt.Printf("\nAudios disponibles de tipo '%s':\n", tipo)

	for _, audio := range respuesta.VectorAudiosPorTipo {
		fmt.Printf("- %s\n", audio.Titulo)
	}

	fmt.Println()
}

/*
Mostrar un menu de tipos de audio que hay disponibles
*/
func MostrarMenuDeTipos(client pb.AudioServiceClient, ctx context.Context) {
	bandera := true
	for bandera {
		fmt.Print("\n Menu de tipos \n")
		fmt.Print("\n 1. Canciones \n")
		fmt.Print("\n 2. Audiolibros\n")
		fmt.Print("\n 3. Ruido Blanco\n")
		fmt.Print("\n 4. Podcast\n")
		fmt.Print("\n 5. Atras\n")

		readerInput := bufio.NewReader(os.Stdin)
		fmt.Print("\nIngrese la opcion del menu: ")
		opcion, _ := readerInput.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			MostrarAudiosPorTipo("canciones", ctx)
		case "2":
			MostrarAudiosPorTipo("Audiolibros", ctx)
		case "3":
			MostrarAudiosPorTipo("Ruido Blanco", ctx)
		case "4":
			MostrarAudiosPorTipo("Podcast", ctx)
		case "5":
			bandera = false
		default:
			fmt.Print("\nOpcion no valida, intentalo nuevamente\n")
		}
	}
}

func MostrarMenuPrincipal(client pb.AudioServiceClient, ctx context.Context) {

	bandera := true
	for bandera == true {
		fmt.Print("\n Menu de cliente \n")
		fmt.Print("\n 1. Ver Tipos de Audio \n")
		fmt.Print("\n 2. Salir\n")

		readerInput := bufio.NewReader(os.Stdin)
		fmt.Print("\nIngrese la opcion del menu: ")
		opcion, _ := readerInput.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			MostrarMenuDeTipos(client, ctx)
		case "2":
			bandera = false
			fmt.Print("\n Adiooooooooos :D\n")
		default:
			fmt.Print("\nOpcion no valida, intentalo nuevamente\n")
		}
	}

}
