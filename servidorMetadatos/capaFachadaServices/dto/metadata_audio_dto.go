package dto

// MetadataAudioDTO es el objeto de transferencia de datos que se recibe al
// registrar un audio y que se devuelve dentro de la respuesta de consulta.

type MetadataAudioDTO struct {
	Titulo     string `json:"titulo"`
	Duracion   int    `json:"duracion"`
	Tipo       string `json:"tipo"`
	Disponible bool   `json:"disponible"`
}

// RespuestaMetadataAudioDTO es el DTO de respuesta para la consulta de un
// audio: incluye el audio encontrado (si aplica), un código de resultado y
// un mensaje descriptivo.

type RespuestaMetadataAudioDTO struct {
	ObjAudio MetadataAudioDTO `json:"objAudio"`
	Codigo   int              `json:"Codigo"`
	Mensaje  string           `json:"Mensaje"`
}
