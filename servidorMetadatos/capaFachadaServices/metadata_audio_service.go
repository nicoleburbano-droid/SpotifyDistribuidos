package service

import (
	"microservicio/capaAccesoADatos/entity"
	"microservicio/capaAccesoADatos/repository"
	"microservicio/capaFachadaServices/dto"
)

// MetadataAudioService es la fachada (facade) que expone al controlador las
// operaciones de negocio, ocultando el acceso al repositorio y la
// conversión entre Entity y DTO.

type MetadataAudioService struct {
	repository *repository.MetadataAudioRepository
}

func NewMetadataAudioService(repository *repository.MetadataAudioRepository) *MetadataAudioService {
	return &MetadataAudioService{repository: repository}
}

// RegistrarAudio recibe un DTO, lo convierte a Entity y lo registra en el
// repositorio.

func (this *MetadataAudioService) RegistrarAudio(audioDTO dto.MetadataAudioDTO) {
	var audio entity.MetadataAudio

	audio.SetTitulo(audioDTO.Titulo)
	audio.SetDuracion(audioDTO.Duracion)
	audio.SetTipo(audioDTO.Tipo)
	audio.SetDisponible(audioDTO.Disponible)

	this.repository.RegistrarAudio(audio)
}

// ConsultarAudio recibe un título, busca el Entity en el repositorio y lo
// convierte a RespuestaMetadataAudioDTO con el código y mensaje según el
// resultado de la búsqueda.

func (this *MetadataAudioService) ConsultarAudio(titulo string) dto.RespuestaMetadataAudioDTO {
	var respuesta dto.RespuestaMetadataAudioDTO

	audio, encontrado := this.repository.BuscarAudio(titulo)

	if encontrado {

		var audioDTO dto.MetadataAudioDTO
		audioDTO.Titulo = audio.GetTitulo()
		audioDTO.Duracion = audio.GetDuracion()
		audioDTO.Tipo = audio.GetTipo()
		audioDTO.Disponible = audio.GetDisponible()

		respuesta.ObjAudio = audioDTO
		respuesta.Codigo = 200
		respuesta.Mensaje = "Metadata del audio encontrada"
	} else {
		respuesta.Codigo = 400
		respuesta.Mensaje = "La metadata del audio no se encontro"
	}

	return respuesta
}

//Obtener AudiosPorTipo retorna todo el vector de audios almacenados
func (this *MetadataAudioService) ObtenerAudiosPorTipo(tipo string) dto.RespuestaAudiosPorTipoDTO {
	var respuesta dto.RespuestaAudiosPorTipoDTO
	vecAudios := this.repository.ObtenerAudiosPorTipo(tipo)
	
	if len(vecAudios) > 0 {
		// Creamos un nuevo slice del tipo DTO con la misma capacidad
		vectorDTOs := make([]dto.MetadataAudioDTO, 0, len(vecAudios))
		
		// Mapeamos cada entidad a su DTO correspondiente
		for _, audio := range vecAudios {
			dtoAudio := dto.MetadataAudioDTO{
				Titulo:     audio.GetTitulo(), 
				Duracion:   audio.GetDuracion(),
				Tipo:       audio.GetTipo(),
				Disponible: audio.GetDisponible(),
			}
			vectorDTOs = append(vectorDTOs, dtoAudio)
		}
		
		// Asignamos el nuevo slice de DTOs a la respuesta
		respuesta.VectorAudiosPorTipo = vectorDTOs
		respuesta.Codigo = 200
		respuesta.Mensaje = "Audios del tipo buscado encontrados"
	} else {
		respuesta.Codigo = 400
		respuesta.Mensaje = "No se encontraron audios del tipo buscado"
	}

	return respuesta
}
