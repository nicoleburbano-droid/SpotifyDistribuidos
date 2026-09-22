package vistas

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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
Fachada para pedirle al servidor de metadatos los audios disponibles para un determinado tipo
*/
func MostrarAudiosPorTipo(tipo string, client pb.AudioServiceClient, ctx context.Context) {

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
			MostrarAudiosPorTipo(ti)
		case "2":
		case "3":
		case "4":
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
