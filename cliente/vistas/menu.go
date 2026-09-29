package vistas

import (
	"bufio"
	"context"

	"fmt"
	"io"
	"log"

	"os"
	"strconv"
	"strings"

	"cliente.local/grpc-cliente/utilidades"

	util "cliente.local/grpc-cliente/utilidades"
	pb "servidorStreaming.local/grpc-servidor/serviciosAudio"
)

/*
Funcion que lee el titulo de un audio que se quiere reproducir
e invoca un procedimiento remoto que se contecta al servidor de audios
*/
func RecibirAudioAReproducir(client pb.AudioServiceClient, ctx context.Context, titulo string) {

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
func MostrarAudiosPorTipo(tipo string, ctx context.Context, client pb.AudioServiceClient) {
	VectorAudiosPorTipo := utilidades.SolicitarMetadata(tipo, ctx)
	
	// Menu de audios disponibles segun el tipo seleccionado
	bandera := true
	for bandera {
	// Mostrar únicamente los títulos.
		fmt.Printf("\nAudios disponibles de tipo '%s':\n", tipo)

		i := 0
		for _, audio := range VectorAudiosPorTipo {
			fmt.Printf("%d - %s\n", i+1, audio.Titulo)
			i++
		}
		fmt.Printf("%d - Vovler", i+1)

		// leer de que audio se quiere mostrar metadatos
		readerInput := bufio.NewReader(os.Stdin)
		fmt.Print("\nIngrese la opcion del menu: ")
		opcion, _ := readerInput.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		numero, err := strconv.Atoi(opcion)
		if err != nil {
			fmt.Println("Error al convertir:", err)
			return
		}

		// mostrar los metadatos del audio seleccionado
		if (numero > len(VectorAudiosPorTipo)+1 || numero < 1){
			fmt.Println("La opcion ingresada no es valida")
		}else if (numero == len(VectorAudiosPorTipo)+1){
			bandera = false
		}else{
			fmt.Printf("Recurso: %s",VectorAudiosPorTipo[numero-1].Titulo)
			fmt.Printf("\nTitulo: %s", VectorAudiosPorTipo[numero-1].Titulo)
			fmt.Printf("\nDuracion: %d", VectorAudiosPorTipo[numero-1].Duracion)
			fmt.Printf("\nTipo: %s \n", VectorAudiosPorTipo[numero-1].Tipo)

			// opcion de reproducir, si o volver. Y en opcion de reproducir ya la otra opcion
			fmt.Printf("\n¿Deseas reproducir el audio?")
			fmt.Printf("\n1. Si")
			fmt.Printf("\n2. No, regresar")
			reproducir, _ := readerInput.ReadString('\n')
			reproducir = strings.TrimSpace(opcion)

			switch reproducir {
			case "1":
				RecibirAudioAReproducir(client, ctx, VectorAudiosPorTipo[numero-1].Titulo)
			case "2":
				bandera = false
			default:
				fmt.Printf("La opcion seleccionada no es valida")
			}
			
		}
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
		fmt.Print("\n 1. Musica \n")
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
			MostrarAudiosPorTipo("Musica", ctx, client)
		case "2":
			MostrarAudiosPorTipo("Audiolibros", ctx, client)
		case "3":
			MostrarAudiosPorTipo("Ruido Blanco", ctx, client)
		case "4":
			MostrarAudiosPorTipo("Podcast", ctx, client)
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
