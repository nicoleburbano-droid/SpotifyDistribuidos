package controller

import (
	"net/http"

	service "microservicio/capaFachadaServices"
	"microservicio/capaFachadaServices/dto"

	"github.com/gin-gonic/gin"
)

// MetadataAudioController expone los servicios REST de audio usando Gin.

type MetadataAudioController struct {
	service *service.MetadataAudioService
}

func NewMetadataAudioController(service *service.MetadataAudioService) *MetadataAudioController {
	return &MetadataAudioController{service: service}
}

// RegistrarAudio - POST /audios
// Recibe un MetadataAudioDTO en el body y lo registra a través de la
// fachada de servicios.
func (this *MetadataAudioController) RegistrarAudio(ctx *gin.Context) {
	var audioDTO dto.MetadataAudioDTO

	if err := ctx.ShouldBindJSON(&audioDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"codigo":  http.StatusBadRequest,
			"mensaje": "Los datos del audio son inválidos: " + err.Error(),
		})
		return
	}

	this.service.RegistrarAudio(audioDTO)

	ctx.JSON(http.StatusCreated, gin.H{
		"codigo":   http.StatusCreated,
		"mensaje":  "Audio registrado correctamente",
		"objAudio": audioDTO,
	})

}

// ConsultarAudio - GET /audios/:titulo
// Consulta un audio por su título y responde con el código HTTP definido
// por la fachada (200 si se encontró, 400 si no).

func (this *MetadataAudioController) ConsultarAudio(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	respuesta := this.service.ConsultarAudio(titulo)
	ctx.JSON(respuesta.Codigo, respuesta)
}


// ObtenerTodos - GET /audios/:tipo
// Consulta todos los audios y responde con el código HTTP definido
// por la fachada (200 si se encontró, 401 si no)

func (this *MetadataAudioController) ConsultarTodos(ctx *gin.Context){
	tipo := ctx.Param("tipo")
	respuesta := this.service.ObtenerAudiosPorTipo(tipo)
	ctx.JSON(respuesta.Codigo, respuesta)
}