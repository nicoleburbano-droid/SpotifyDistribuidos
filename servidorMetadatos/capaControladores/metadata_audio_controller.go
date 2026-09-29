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

// TODO registrarMusica
func (this *MetadataAudioController) RegistrarMusica(ctx *gin.Context) {
	var musicaDTO dto.MetadataMusicaDTO
	if err := ctx.ShouldBindJSON(&musicaDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"codigo": http.StatusBadRequest, "mensaje": "Los datos de la música son inválidos: " + err.Error()})
		return
	}
	this.service.RegistrarMusica(musicaDTO)
	ctx.JSON(http.StatusCreated, gin.H{"codigo": http.StatusCreated, "mensaje": "Música registrada correctamente", "objMusica": musicaDTO})
}

// TODO registrarAudioLibro
func (this *MetadataAudioController) RegistrarAudioLibro(ctx *gin.Context) {
	var libroDTO dto.MetadataAudioLibrosDTO
	if err := ctx.ShouldBindJSON(&libroDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"codigo": http.StatusBadRequest, "mensaje": "Los datos del audiolibro son inválidos: " + err.Error()})
		return
	}
	this.service.RegistrarAudioLibro(libroDTO)
	ctx.JSON(http.StatusCreated, gin.H{"codigo": http.StatusCreated, "mensaje": "Audiolibro registrado correctamente", "objAudioLibro": libroDTO})
}

// TODO registrarPodcast
func (this *MetadataAudioController) RegistrarPodcast(ctx *gin.Context) {
	var podcastDTO dto.MetadataPodcastDTO
	if err := ctx.ShouldBindJSON(&podcastDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"codigo": http.StatusBadRequest, "mensaje": "Los datos del podcast son inválidos: " + err.Error()})
		return
	}
	this.service.RegistrarPodcast(podcastDTO)
	ctx.JSON(http.StatusCreated, gin.H{"codigo": http.StatusCreated, "mensaje": "Podcast registrado correctamente", "objPodcast": podcastDTO})
}

// TODO registrarRuidoBlanco
func (this *MetadataAudioController) RegistrarRuidoBlanco(ctx *gin.Context) {
	var ruidoDTO dto.MetadataRuidoBlancoDTO
	if err := ctx.ShouldBindJSON(&ruidoDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"codigo": http.StatusBadRequest, "mensaje": "Los datos del ruido blanco son inválidos: " + err.Error()})
		return
	}
	this.service.RegistrarRuidoBlanco(ruidoDTO)
	ctx.JSON(http.StatusCreated, gin.H{"codigo": http.StatusCreated, "mensaje": "Ruido blanco registrado correctamente", "objRuidoBlanco": ruidoDTO})
}

// ConsultarAudio - GET /audios/:titulo
// Consulta un audio por su título y responde con el código HTTP definido
// por la fachada (200 si se encontró, 400 si no).

func (this *MetadataAudioController) ConsultarAudio(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	respuesta := this.service.ConsultarAudio(titulo)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ConsultarMusica
func (this *MetadataAudioController) ConsultarMusica(ctx *gin.Context) {
	respuesta := this.service.ConsultarMusica(ctx.Param("titulo"))
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ConsultarAudioLibro
func (this *MetadataAudioController) ConsultarAudioLibro(ctx *gin.Context) {
	respuesta := this.service.ConsultarAudioLibro(ctx.Param("titulo"))
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ConsultarPodcast
func (this *MetadataAudioController) ConsultarPodcast(ctx *gin.Context) {
	respuesta := this.service.ConsultarPodcast(ctx.Param("titulo"))
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ConsultarRuidoBlanco
func (this *MetadataAudioController) ConsultarRuidoBlanco(ctx *gin.Context) {
	respuesta := this.service.ConsultarRuidoBlanco(ctx.Param("tipoSonido"))
	ctx.JSON(respuesta.Codigo, respuesta)
}

// ObtenerTodos - GET /audios/:tipo
// Consulta todos los audios y responde con el código HTTP definido
// por la fachada (200 si se encontró, 401 si no)

func (this *MetadataAudioController) ConsultarAudiosPorTipo(ctx *gin.Context) {
	tipo := ctx.Param("tipo")
	respuesta := this.service.ObtenerAudiosPorTipo(tipo)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ObtenerMusica
func (this *MetadataAudioController) ObtenerMusica(ctx *gin.Context) {
	respuesta := this.service.ObtenerMusica()
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ObtenerAudioLibros
func (this *MetadataAudioController) ObtenerAudioLibros(ctx *gin.Context) {
	respuesta := this.service.ObtenerAudioLibros()
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ObtenerPodcast
func (this *MetadataAudioController) ObtenerPodcast(ctx *gin.Context) {
	respuesta := this.service.ObtenerPodcast()
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ObtenerRuidoBlanco
func (this *MetadataAudioController) ObtenerRuidoBlanco(ctx *gin.Context) {
	respuesta := this.service.ObtenerRuidoBlanco()
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ObtenerTiposDeAudio
func (this *MetadataAudioController) ObtenerTiposDeAudio(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, this.service.ObtenerTiposDeAudio())
}
