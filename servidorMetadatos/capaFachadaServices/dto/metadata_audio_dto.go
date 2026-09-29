package dto

// MetadataAudioDTO es el objeto de transferencia de datos que se recibe al
// registrar un audio y que se devuelve dentro de la respuesta de consulta.

type MetadataAudioDTO struct {
	Titulo     string `json:"titulo"`
	Duracion   int    `json:"duracion"`
	Tipo       string `json:"tipo"`
	Disponible bool   `json:"disponible"`
}


type MetadataMusicaDTO struct {
	ArtistaPrincipal  string `json:"artistaPrincipal"`
	Album             string `json:"album"`
	Genero            string `json:"genero"`
	TituloCancion     string `json:"tituloCancion"`
	SelloDiscografico string `json:"selloDiscografico"`
	AnioLanzamiento   string `json:"anioLanzamiento"`
}

type MetadataAudioLibrosDTO struct {
	TituloLibro string `json:"tituloLibro"`
	Autor       string `json:"autor"`
	Narrador    string `json:"narrador"`
	Editorial   string `json:"editorial"`
	Isbn        int    `json:"isbn"`
	Capitulo    int    `json:"capitulo"`
}

type MetadataPodcastDTO struct {
	NombrePodcast          string `json:"nombrePodcast"`
	TituloEpisodio         string `json:"tituloEpisodio"`
	Anfitrion              string `json:"anfitrion"`
	NumeroTemporada        int    `json:"numeroTemporada"`
	NotasShow              string `json:"notasShow"`
	ClasificacionContenido string `json:"clasificacionContenido"`
}

type MetadataRuidoBlancoDTO struct {
	TipoSonido          string `json:"tipoSonido"`
	FuenteAudio         string `json:"fuenteAudio"`
	UsoSugerido         string `json:"usoSugerido"`
	ProveedorContenido  string `json:"proveedorContenido"`
	DuracionBucle       int    `json:"duracionBucle"`
	FrecuenciaDominante string `json:"frecuenciaDominante"`
}

type TipoAudioDTO struct {
	Id   int    `json:"id"`
	Tipo string `json:"tipo"`
}


// RespuestaMetadataAudioDTO es el DTO de respuesta para la consulta de un
// audio: incluye el audio encontrado (si aplica), un código de resultado y
// un mensaje descriptivo.

type RespuestaMetadataAudioDTO struct {
	ObjAudio MetadataAudioDTO `json:"objAudio"`
	Codigo   int              `json:"Codigo"`
	Mensaje  string           `json:"Mensaje"`
}

// RespuestaAudiosPorTipoDTO es el DTO para obtener todos los metadados de los audios que
// pertenecen a un determinado tipo
type RespuestaAudiosPorTipoDTO struct {
	VectorAudiosPorTipo []MetadataAudioDTO	`json:"vectorAudiosPorTipo"`
	Codigo 				int 				`json:"Codigo"`
	Mensaje 			string 				`json:"Mensaje"`
}

type RespuestaMusicaDTO struct {
	VectorMusica []MetadataMusicaDTO	    `json:"vectorMusica"`
	Codigo 				int 				`json:"Codigo"`
	Mensaje 			string 				`json:"Mensaje"`
}

type RespuestaAudioLibrosDTO struct {
	VectorAudioLibros []MetadataAudioLibrosDTO	`json:"vectorAudioLibros"`
	Codigo 				int 				`json:"Codigo"`
	Mensaje 			string 				`json:"Mensaje"`
}

type RespuestaPodcastDTO struct {
	VectorPodcast []MetadataPodcastDTO	`json:"vectorPodcast"`
	Codigo 				int 				`json:"Codigo"`
	Mensaje 			string 				`json:"Mensaje"`
}

type RespuestaRuidoBlancoDTO struct {
	VectorRuidoBlanco []MetadataRuidoBlancoDTO	`json:"vectorRuidoBlanco"`
	Codigo 				int 				`json:"Codigo"`
	Mensaje 			string 				`json:"Mensaje"`
}