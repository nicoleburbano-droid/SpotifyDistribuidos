package controladores

import (
	"administrador/dtos"
	"administrador/servicios"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ControladorAdministrador conecta la vista con el servicio REST.
type ControladorAdministrador struct {
	servicio *servicios.ServicioAudios
}

// NuevoControladorAdministrador crea el controlador.
func NuevoControladorAdministrador(servicio *servicios.ServicioAudios) *ControladorAdministrador {
	return &ControladorAdministrador{servicio: servicio}
}

// ValidarArchivo verifica que el archivo exista y sea .mp3 antes de pedir
// los metadatos, para no llenar todo el formulario en vano.
func (this *ControladorAdministrador) ValidarArchivo(rutaArchivo string) error {
	info, err := os.Stat(rutaArchivo)
	if err != nil || info.IsDir() {
		return fmt.Errorf("el archivo %s no existe", rutaArchivo)
	}
	if strings.ToLower(filepath.Ext(rutaArchivo)) != ".mp3" {
		return fmt.Errorf("el archivo debe ser .mp3")
	}
	return nil
}

// AlmacenarMusica sube una canción.
func (this *ControladorAdministrador) AlmacenarMusica(rutaArchivo string, musica dtos.MetadataMusicaDTO) (dtos.RespuestaAlmacenamientoDTO, error) {
	return this.servicio.AlmacenarAudio(rutaArchivo, "musica", musica)
}

// AlmacenarAudioLibro sube un audiolibro.
func (this *ControladorAdministrador) AlmacenarAudioLibro(rutaArchivo string, libro dtos.MetadataAudioLibrosDTO) (dtos.RespuestaAlmacenamientoDTO, error) {
	return this.servicio.AlmacenarAudio(rutaArchivo, "audiolibro", libro)
}

// AlmacenarPodcast sube un podcast.
func (this *ControladorAdministrador) AlmacenarPodcast(rutaArchivo string, podcast dtos.MetadataPodcastDTO) (dtos.RespuestaAlmacenamientoDTO, error) {
	return this.servicio.AlmacenarAudio(rutaArchivo, "podcast", podcast)
}

// AlmacenarRuidoBlanco sube un ruido blanco.
func (this *ControladorAdministrador) AlmacenarRuidoBlanco(rutaArchivo string, ruido dtos.MetadataRuidoBlancoDTO) (dtos.RespuestaAlmacenamientoDTO, error) {
	return this.servicio.AlmacenarAudio(rutaArchivo, "ruidoBlanco", ruido)
}
