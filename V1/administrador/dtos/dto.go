package dtos

// Los DTOs de metadatos usan los mismos nombres JSON que el servidor de
// metadatos, para que el servidor de audios los pueda reenviar tal cual.

// MetadataMusicaDTO transporta los metadatos de una canción.
type MetadataMusicaDTO struct {
	ArtistaPrincipal  string `json:"artistaPrincipal"`
	Album             string `json:"album"`
	Genero            string `json:"genero"`
	TituloCancion     string `json:"tituloCancion"`
	SelloDiscografico string `json:"selloDiscografico"`
	AnioLanzamiento   string `json:"anioLanzamiento"`
}

// MetadataAudioLibrosDTO transporta los metadatos de un audiolibro.
type MetadataAudioLibrosDTO struct {
	TituloLibro string `json:"tituloLibro"`
	Autor       string `json:"autor"`
	Narrador    string `json:"narrador"`
	Editorial   string `json:"editorial"`
	Isbn        int    `json:"isbn"`
	Capitulo    int    `json:"capitulo"`
}

// MetadataPodcastDTO transporta los metadatos de un podcast.
type MetadataPodcastDTO struct {
	NombrePodcast          string `json:"nombrePodcast"`
	TituloEpisodio         string `json:"tituloEpisodio"`
	Anfitrion              string `json:"anfitrion"`
	NumeroTemporada        int    `json:"numeroTemporada"`
	NotasShow              string `json:"notasShow"`
	ClasificacionContenido string `json:"clasificacionContenido"`
}

// MetadataRuidoBlancoDTO transporta los metadatos de un ruido blanco.
type MetadataRuidoBlancoDTO struct {
	TipoSonido          string `json:"tipoSonido"`
	FuenteAudio         string `json:"fuenteAudio"`
	UsoSugerido         string `json:"usoSugerido"`
	ProveedorContenido  string `json:"proveedorContenido"`
	DuracionBucle       int    `json:"duracionBucle"`
	FrecuenciaDominante string `json:"frecuenciaDominante"`
}

// RespuestaAlmacenamientoDTO es la respuesta del servidor de audios.
type RespuestaAlmacenamientoDTO struct {
	Codigo        int    `json:"codigo"`
	Mensaje       string `json:"mensaje"`
	NombreArchivo string `json:"nombreArchivo"`
}
