package vistas

import (
	"administrador/controladores"
	"administrador/dtos"
	"fmt"
)

// MostrarMenuPrincipal muestra el menú del administrador.
func MostrarMenuPrincipal(controlador *controladores.ControladorAdministrador) {
	for {
		fmt.Print("\n Menu de administrador \n")
		fmt.Print("\n 1. Almacenar audio \n")
		fmt.Print("\n 2. Salir\n")

		switch leerOpcion(2) {
		case 1:
			mostrarMenuDeTipos(controlador)
		case 2:
			fmt.Print("\n Adiooooooooos :D\n")
			return
		}
	}
}

// mostrarMenuDeTipos permite elegir el tipo de audio a almacenar. El orden
// es el mismo del menú del cliente.
func mostrarMenuDeTipos(controlador *controladores.ControladorAdministrador) {
	fmt.Print("\n Tipo de audio a almacenar \n")
	fmt.Print("\n 1. Musica \n")
	fmt.Print("\n 2. Audiolibros\n")
	fmt.Print("\n 3. Podcast\n")
	fmt.Print("\n 4. Ruido Blanco\n")
	fmt.Print("\n 5. Atras\n")

	opcion := leerOpcion(5)
	if opcion == 0 || opcion == 5 {
		return
	}

	rutaArchivo := pedirRutaArchivo(controlador)

	var respuesta dtos.RespuestaAlmacenamientoDTO
	var err error
	switch opcion {
	case 1:
		respuesta, err = controlador.AlmacenarMusica(rutaArchivo, pedirMetadatosMusica())
	case 2:
		respuesta, err = controlador.AlmacenarAudioLibro(rutaArchivo, pedirMetadatosAudioLibro())
	case 3:
		respuesta, err = controlador.AlmacenarPodcast(rutaArchivo, pedirMetadatosPodcast())
	case 4:
		respuesta, err = controlador.AlmacenarRuidoBlanco(rutaArchivo, pedirMetadatosRuidoBlanco())
	}

	if err != nil {
		fmt.Println("\nNo se pudo almacenar el audio:", err)
		return
	}
	fmt.Printf("\n%s\n", respuesta.Mensaje)
}

// pedirRutaArchivo pide la ruta del mp3 hasta que sea válida.
func pedirRutaArchivo(controlador *controladores.ControladorAdministrador) string {
	for {
		rutaArchivo := pedirTexto("\nRuta del archivo .mp3")
		if err := controlador.ValidarArchivo(rutaArchivo); err != nil {
			fmt.Println("Error:", err)
			continue
		}
		return rutaArchivo
	}
}

func pedirMetadatosMusica() dtos.MetadataMusicaDTO {
	return dtos.MetadataMusicaDTO{
		TituloCancion:     pedirTexto("Titulo de la cancion"),
		ArtistaPrincipal:  pedirTexto("Artista principal"),
		Album:             pedirTexto("Album"),
		Genero:            pedirTexto("Genero"),
		SelloDiscografico: pedirTexto("Sello discografico"),
		AnioLanzamiento:   pedirTexto("Año de lanzamiento"),
	}
}

func pedirMetadatosAudioLibro() dtos.MetadataAudioLibrosDTO {
	return dtos.MetadataAudioLibrosDTO{
		TituloLibro: pedirTexto("Titulo del libro"),
		Autor:       pedirTexto("Autor"),
		Narrador:    pedirTexto("Narrador"),
		Editorial:   pedirTexto("Editorial"),
		Isbn:        pedirEntero("ISBN (solo numeros)"),
		Capitulo:    pedirEntero("Capitulo"),
	}
}

func pedirMetadatosPodcast() dtos.MetadataPodcastDTO {
	return dtos.MetadataPodcastDTO{
		NombrePodcast:          pedirTexto("Nombre del podcast"),
		TituloEpisodio:         pedirTexto("Titulo del episodio"),
		Anfitrion:              pedirTexto("Anfitrion"),
		NumeroTemporada:        pedirEntero("Numero de temporada"),
		NotasShow:              pedirTexto("Notas del show"),
		ClasificacionContenido: pedirTexto("Clasificacion (Explicito / Para toda la familia)"),
	}
}

func pedirMetadatosRuidoBlanco() dtos.MetadataRuidoBlancoDTO {
	return dtos.MetadataRuidoBlancoDTO{
		TipoSonido:          pedirTexto("Tipo de sonido"),
		FuenteAudio:         pedirTexto("Fuente del audio"),
		UsoSugerido:         pedirTexto("Uso sugerido"),
		ProveedorContenido:  pedirTexto("Proveedor de contenido"),
		DuracionBucle:       pedirEntero("Duracion del bucle (segundos)"),
		FrecuenciaDominante: pedirTexto("Frecuencia dominante (Graves / Agudos)"),
	}
}
