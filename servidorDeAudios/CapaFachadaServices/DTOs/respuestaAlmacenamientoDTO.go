package dtos

// RespuestaAlmacenamientoDTO es la respuesta JSON que recibe el administrador.
type RespuestaAlmacenamientoDTO struct {
	Codigo        int    `json:"codigo"`
	Mensaje       string `json:"mensaje"`
	NombreArchivo string `json:"nombreArchivo,omitempty"`
}
