package componenteclientemetadatos

import (
	"fmt"

	dto "almacenamiento/CapaFachadaServices/DTOs"
)

// ClienteMetadatos coordina el registro de los metadatos con el servicio local.
type ClienteMetadatos struct {
	serviceMetadatos *MetadataAudioService
}

// NuevoClienteMetadatos crea el cliente del servicio de metadatos.
func NuevoClienteMetadatos(serviceMetadatos *MetadataAudioService) *ClienteMetadatos {
	return &ClienteMetadatos{serviceMetadatos: serviceMetadatos}
}

// RegistrarMetadatos envía al servicio el DTO correspondiente al tipo de audio.
func (this *ClienteMetadatos) RegistrarMetadatos(tipo string, metadatosDTO interface{}) error {
	switch tipo {
	case "musica":
		musica, ok := metadatosDTO.(dto.MetadataMusicaDTO)
		if !ok {
			return fmt.Errorf("los metadatos para %q deben ser MetadataMusicaDTO (recibido %T)", tipo, metadatosDTO)
		}
		this.serviceMetadatos.RegistrarMusica(musica)
	case "audiolibro", "audioLibros":
		libro, ok := metadatosDTO.(dto.MetadataAudioLibrosDTO)
		if !ok {
			return fmt.Errorf("los metadatos para %q deben ser MetadataAudioLibrosDTO (recibido %T)", tipo, metadatosDTO)
		}
		this.serviceMetadatos.RegistrarAudioLibro(libro)
	case "podcast":
		podcast, ok := metadatosDTO.(dto.MetadataPodcastDTO)
		if !ok {
			return fmt.Errorf("los metadatos para %q deben ser MetadataPodcastDTO (recibido %T)", tipo, metadatosDTO)
		}
		this.serviceMetadatos.RegistrarPodcast(podcast)
	case "ruidoBlanco":
		ruido, ok := metadatosDTO.(dto.MetadataRuidoBlancoDTO)
		if !ok {
			return fmt.Errorf("los metadatos para %q deben ser MetadataRuidoBlancoDTO (recibido %T)", tipo, metadatosDTO)
		}
		this.serviceMetadatos.RegistrarRuidoBlanco(ruido)
	default:
		return fmt.Errorf("el tipo %q no existe", tipo)
	}

	return nil
}
