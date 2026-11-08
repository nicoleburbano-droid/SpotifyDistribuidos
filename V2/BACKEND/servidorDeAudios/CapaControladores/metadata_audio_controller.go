package capacontroladores

import (
	"fmt"
	"net/http"

	service "almacenamiento/CapaFachadaServices/ComponenteClienteMetadatos"
	

	"github.com/gin-gonic/gin"
)

// MetadataAudioController expone los servicios REST de audio usando Gin.

// Definicion de la clase, tiene un atributo de tipo service, que se inyecta por el constructor
type MetadataAudioController struct {
	service *service.MetadataAudioService
}

//constructor de la calse metadaAudioController
func NewMetadataAudioController(service *service.MetadataAudioService) *MetadataAudioController {
	return &MetadataAudioController{service: service}
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

//  ConsultarMusica
func (this *MetadataAudioController) ConsultarMusica(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	fmt.Printf("[REST] GET /musica/titulo/%s: consulta de música\n", titulo)
	respuesta := this.service.ConsultarMusica(titulo)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// ConsultarAudioLibro
func (this *MetadataAudioController) ConsultarAudioLibro(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	fmt.Printf("[REST] GET /audiolibros/titulo/%s: consulta de audiolibro\n", titulo)
	respuesta := this.service.ConsultarAudioLibro(titulo)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// ConsultarPodcast
func (this *MetadataAudioController) ConsultarPodcast(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	fmt.Printf("[REST] GET /podcasts/titulo/%s: consulta de podcast\n", titulo)
	respuesta := this.service.ConsultarPodcast(titulo)
	ctx.JSON(respuesta.Codigo, respuesta)
}

// ConsultarRuidoBlanco
func (this *MetadataAudioController) ConsultarRuidoBlanco(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	fmt.Printf("[REST] GET /ruido-blanco/titulo/%s: consulta de ruido blanco\n", titulo)
	respuesta := this.service.ConsultarRuidoBlanco(titulo)
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

// ObtenerMusica
func (this *MetadataAudioController) ObtenerMusica(ctx *gin.Context) {
	fmt.Printf("[REST] GET /musica/all: consulta de toda la música\n")
	respuesta := this.service.ObtenerMusica()
	ctx.JSON(respuesta.Codigo, respuesta)
}

// ObtenerAudioLibros
func (this *MetadataAudioController) ObtenerAudioLibros(ctx *gin.Context) {
	fmt.Printf("[REST] GET /audiolibros/all: consulta de todos los audiolibros\n")
	respuesta := this.service.ObtenerAudioLibros()
	ctx.JSON(respuesta.Codigo, respuesta)
}

// ObtenerPodcast
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
