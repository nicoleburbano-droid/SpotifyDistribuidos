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
func MostrarMusica(ctx context.Context, client pb.AudioServiceClient) {

	VectorAudios := util.SolicitarMusica(ctx)

	// Menu de audios disponibles
	bandera := true
	for bandera {
		// Mostrar únicamente los títulos.
		fmt.Printf("\nAudios disponibles de Musica:\n")

		i := 0
		for _, audio := range VectorAudios {
			fmt.Printf("%d - %s\n", i+1, audio.TituloCancion)
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
		if numero > len(VectorAudios)+1 || numero < 1 {
			fmt.Println("La opcion ingresada no es valida")
		} else if numero == len(VectorAudios)+1 {
			bandera = false
		} else {
			fmt.Printf("\nRecurso: %s\n", VectorAudios[numero-1].TituloCancion)
			fmt.Printf("\nTitulo: %s", VectorAudios[numero-1].TituloCancion)
			fmt.Printf("\nAlbum: %s", VectorAudios[numero-1].Album)
			fmt.Printf("\nAño Lanzamiento: %s", VectorAudios[numero-1].AnioLanzamiento)
			fmt.Printf("\nArtista Principal: %s", VectorAudios[numero-1].ArtistaPrincipal)
			fmt.Printf("\nGenero: %s", VectorAudios[numero-1].Genero)
			fmt.Printf("\nSello Discografico: %s \n", VectorAudios[numero-1].SelloDiscografico)

			// opcion de reproducir, si o volver. Y en opcion de reproducir ya la otra opcion
			fmt.Printf("\n¿Deseas reproducir el audio?")
			fmt.Printf("\n1. Si")
			fmt.Printf("\n2. No, regresar")
			fmt.Println("")
			reproducir, _ := readerInput.ReadString('\n')
			reproducir = strings.TrimSpace(reproducir)

			switch reproducir {
			case "1":
				RecibirAudioAReproducir(client, ctx, VectorAudios[numero-1].TituloCancion+".mp3")
			case "2":
				bandera = false
			default:
				fmt.Printf("La opcion seleccionada no es valida")
			}
		}
	}

	fmt.Println()
}

func MostrarAudiolibros(ctx context.Context, client pb.AudioServiceClient) {

	VectorAudios := util.SolicitarAudioLibros(ctx)

	// Menu de audios disponibles
	bandera := true
	for bandera {
		// Mostrar únicamente los títulos.
		fmt.Printf("\nAudios disponibles de Audio Libros:\n")

		i := 0
		for _, audio := range VectorAudios {
			fmt.Printf("%d - %s\n", i+1, audio.TituloLibro)
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
		if numero > len(VectorAudios)+1 || numero < 1 {
			fmt.Println("La opcion ingresada no es valida")
		} else if numero == len(VectorAudios)+1 {
			bandera = false
		} else {
			fmt.Printf("\nRecurso: %s \n", VectorAudios[numero-1].TituloLibro)
			fmt.Printf("\nTitulo: %s", VectorAudios[numero-1].TituloLibro)
			fmt.Printf("\nAutor: %s", VectorAudios[numero-1].Autor)
			fmt.Printf("\nEditorial: %s", VectorAudios[numero-1].Editorial)
			fmt.Printf("\nNarrador: %s", VectorAudios[numero-1].Narrador)
			fmt.Printf("\nCapitulo: %d", VectorAudios[numero-1].Capitulo)
			fmt.Printf("\nISBN: %d \n", VectorAudios[numero-1].Isbn)

			// opcion de reproducir, si o volver. Y en opcion de reproducir ya la otra opcion
			fmt.Printf("\n¿Deseas reproducir el audio?")
			fmt.Printf("\n1. Si")
			fmt.Printf("\n2. No, regresar")
			fmt.Println("")
			reproducir, _ := readerInput.ReadString('\n')
			reproducir = strings.TrimSpace(reproducir)

			switch reproducir {
			case "1":
				RecibirAudioAReproducir(client, ctx, VectorAudios[numero-1].TituloLibro+".mp3")
			case "2":
				bandera = false
			default:
				fmt.Printf("La opcion seleccionada no es valida")
			}
		}
	}

	fmt.Println()
}

func MostrarPodcast(ctx context.Context, client pb.AudioServiceClient) {
	VectorAudios := util.SolicitarPodcast(ctx)

	// Menu de audios disponibles
	bandera := true
	for bandera {
		// Mostrar únicamente los títulos.
		fmt.Printf("\nAudios disponibles de Podcast:\n")

		i := 0
		for _, audio := range VectorAudios {
			fmt.Printf("%d - %s\n", i+1, audio.NombrePodcast)
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
		if numero > len(VectorAudios)+1 || numero < 1 {
			fmt.Println("La opcion ingresada no es valida")
		} else if numero == len(VectorAudios)+1 {
			bandera = false
		} else {
			fmt.Printf("\nRecurso: %s\n", VectorAudios[numero-1].NombrePodcast)
			fmt.Printf("\nNombre Podcast: %s", VectorAudios[numero-1].NombrePodcast)
			fmt.Printf("\nAnfitrion: %s", VectorAudios[numero-1].Anfitrion)
			fmt.Printf("\nClasificacion: %s", VectorAudios[numero-1].ClasificacionContenido)
			fmt.Printf("\nNotas del Show: %s", VectorAudios[numero-1].NotasShow)
			fmt.Printf("\nEpisodio: %s", VectorAudios[numero-1].TituloEpisodio)
			fmt.Printf("\nNo. Temporada: %d", VectorAudios[numero-1].NumeroTemporada)

			// opcion de reproducir, si o volver. Y en opcion de reproducir ya la otra opcion
			fmt.Printf("\n¿Deseas reproducir el audio?")
			fmt.Printf("\n1. Si")
			fmt.Printf("\n2. No, regresar")
			fmt.Println("")
			reproducir, _ := readerInput.ReadString('\n')
			reproducir = strings.TrimSpace(reproducir)

			switch reproducir {
			case "1":
				RecibirAudioAReproducir(client, ctx, VectorAudios[numero-1].NombrePodcast+".mp3")
			case "2":
				bandera = false
			default:
				fmt.Printf("La opcion seleccionada no es valida")
			}
		}
	}

	fmt.Println()
}

func MostrarRuidoBlanco(ctx context.Context, client pb.AudioServiceClient) {
	VectorAudios := util.SolicitarRuidoBlanco(ctx)

	// Menu de audios disponibles
	bandera := true
	for bandera {
		// Mostrar únicamente los títulos.
		fmt.Printf("\nAudios disponibles de Ruido Blanco:\n")

		i := 0
		for _, audio := range VectorAudios {
			fmt.Printf("%d - %s\n", i+1, audio.TipoSonido)
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
		if numero > len(VectorAudios)+1 || numero < 1 {
			fmt.Println("La opcion ingresada no es valida")
		} else if numero == len(VectorAudios)+1 {
			bandera = false
		} else {
			fmt.Printf("\nRecurso: %s\n", VectorAudios[numero-1].TipoSonido)
			fmt.Printf("\nSonido: %s", VectorAudios[numero-1].TipoSonido)
			fmt.Printf("\nFrecuencia: %s", VectorAudios[numero-1].FrecuenciaDominante)
			fmt.Printf("\nFuente de Audio: %s", VectorAudios[numero-1].FuenteAudio)
			fmt.Printf("\nProveedor: %s", VectorAudios[numero-1].ProveedorContenido)
			fmt.Printf("\nUso: %s", VectorAudios[numero-1].UsoSugerido)
			fmt.Printf("\nDuracion Bucle: %d", VectorAudios[numero-1].DuracionBucle)

			// opcion de reproducir, si o volver. Y en opcion de reproducir ya la otra opcion
			fmt.Printf("\n¿Deseas reproducir el audio?")
			fmt.Printf("\n1. Si")
			fmt.Printf("\n2. No, regresar")
			fmt.Println("")
			reproducir, _ := readerInput.ReadString('\n')
			reproducir = strings.TrimSpace(reproducir)

			switch reproducir {
			case "1":
				RecibirAudioAReproducir(client, ctx, VectorAudios[numero-1].TipoSonido+".mp3")
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
		fmt.Print("\n 3. Podcast\n")
		fmt.Print("\n 4. Ruido Blanco\n")
		fmt.Print("\n 5. Atras\n")

		readerInput := bufio.NewReader(os.Stdin)
		fmt.Print("\nIngrese la opcion del menu: ")
		opcion, _ := readerInput.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			MostrarMusica(ctx, client)
		case "2":
			MostrarAudiolibros(ctx, client)
		case "3":
			MostrarPodcast(ctx, client)
		case "4":
			MostrarRuidoBlanco(ctx, client)
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
