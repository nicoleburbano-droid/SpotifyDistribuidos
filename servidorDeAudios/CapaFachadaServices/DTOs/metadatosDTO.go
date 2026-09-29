package dtos

// Estos DTOs tienen los mismos nombres JSON que los del servidor de
// metadatos (servidorMetadatos/capaFachadaServices/dto), porque el servidor
// de audios los reenvía tal cual para registrar el audio.

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
