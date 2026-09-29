package controller

import (
	"fmt"
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
	fmt.Printf("[REST] POST /audios: solicitud para registrar audio general\n")
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
	fmt.Printf("[REST] POST /musica: solicitud para registrar música\n")
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
	fmt.Printf("[REST] POST /audiolibros: solicitud para registrar audiolibro\n")
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
	fmt.Printf("[REST] POST /podcasts: solicitud para registrar podcast\n")
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
	fmt.Printf("[REST] POST /ruido-blanco: solicitud para registrar ruido blanco\n")
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
	fmt.Printf("[REST] GET /audios/titulo/%s: consulta de audio\n", titulo)
	respuesta := this.service.ConsultarAudio(titulo)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ConsultarMusica
func (this *MetadataAudioController) ConsultarMusica(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	fmt.Printf("[REST] GET /musica/titulo/%s: consulta de música\n", titulo)
	respuesta := this.service.ConsultarMusica(titulo)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ConsultarAudioLibro
func (this *MetadataAudioController) ConsultarAudioLibro(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	fmt.Printf("[REST] GET /audiolibros/titulo/%s: consulta de audiolibro\n", titulo)
	respuesta := this.service.ConsultarAudioLibro(titulo)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ConsultarPodcast
func (this *MetadataAudioController) ConsultarPodcast(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	fmt.Printf("[REST] GET /podcasts/titulo/%s: consulta de podcast\n", titulo)
	respuesta := this.service.ConsultarPodcast(titulo)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ConsultarRuidoBlanco
func (this *MetadataAudioController) ConsultarRuidoBlanco(ctx *gin.Context) {
	tipoSonido := ctx.Param("tipoSonido")
	fmt.Printf("[REST] GET /ruido-blanco/tipo/%s: consulta de ruido blanco\n", tipoSonido)
	respuesta := this.service.ConsultarRuidoBlanco(tipoSonido)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// ObtenerTodos - GET /audios/:tipo
// Consulta todos los audios y responde con el código HTTP definido
// por la fachada (200 si se encontró, 401 si no)

func (this *MetadataAudioController) ConsultarAudiosPorTipo(ctx *gin.Context) {
	tipo := ctx.Param("tipo")
	fmt.Printf("[REST] GET /audios/tipo/%s: consulta de audios por tipo\n", tipo)
	respuesta := this.service.ObtenerAudiosPorTipo(tipo)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ObtenerMusica
func (this *MetadataAudioController) ObtenerMusica(ctx *gin.Context) {
	fmt.Printf("[REST] GET /musica/all: consulta de toda la música\n")
	respuesta := this.service.ObtenerMusica()
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ObtenerAudioLibros
func (this *MetadataAudioController) ObtenerAudioLibros(ctx *gin.Context) {
	fmt.Printf("[REST] GET /audiolibros/all: consulta de todos los audiolibros\n")
	respuesta := this.service.ObtenerAudioLibros()
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ObtenerPodcast
func (this *MetadataAudioController) ObtenerPodcast(ctx *gin.Context) {
	fmt.Printf("[REST] GET /podcasts/all: consulta de todos los podcasts\n")
	respuesta := this.service.ObtenerPodcast()
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ObtenerRuidoBlanco
func (this *MetadataAudioController) ObtenerRuidoBlanco(ctx *gin.Context) {
	fmt.Printf("[REST] GET /ruido-blanco/all: consulta de todos los ruidos blancos\n")
	respuesta := this.service.ObtenerRuidoBlanco()
	ctx.JSON(respuesta.Codigo, respuesta)
}

// TODO ObtenerTiposDeAudio
func (this *MetadataAudioController) ObtenerTiposDeAudio(ctx *gin.Context) {
	fmt.Printf("[REST] GET /tipos-audio: consulta de tipos de audio disponibles\n")
	ctx.JSON(http.StatusOK, this.service.ObtenerTiposDeAudio())
}
