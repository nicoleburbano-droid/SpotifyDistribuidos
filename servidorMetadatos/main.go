package main

import (
	"microservicio/capaAccesoADatos/repository"
	controller "microservicio/capaControladores"
	service "microservicio/capaFachadaServices"

	"github.com/gin-gonic/gin"
)

func main() {

	audioRepositry := repository.NewMetadataAudioRepository()
	audioService := service.NewMetadataAudioService(audioRepositry)
	audioController := controller.NewMetadataAudioController(audioService)

	router := gin.Default()

	// audios
	router.POST("/audios", audioController.RegistrarAudio)
	router.GET("/audios/titulo/:titulo", audioController.ConsultarAudio)
	router.GET("audios/tipo/:tipo", audioController.ConsultarAudiosPorTipo)

	// Musica

	router.POST("/musica", audioController.RegistrarMusica) 
	router.GET("/musica/titulo/:titulo", audioController.ConsultarMusica)
	router.GET("/musica/all", audioController.ObtenerMusica)

	// Audiolibros

	router.POST("/audiolibros", audioController.RegistrarAudioLibro)
	router.GET("/audiolibros/titulo/:titulo", audioController.ConsultarAudioLibro)
	router.GET("/audiolibros/all", audioController.ObtenerAudioLibros)

	// Podcast

	router.POST("/podcasts", audioController.RegistrarPodcast)
	router.GET("/podcasts/titulo/:titulo", audioController.ConsultarPodcast)
	router.GET("/podcasts/all", audioController.ObtenerPodcast)

	// Ruido blanco

	router.POST("/ruido-blanco", audioController.RegistrarRuidoBlanco)
	router.GET("/ruido-blanco/tipo/:tipoSonido", audioController.ConsultarRuidoBlanco)
	router.GET("/ruido-blanco/all", audioController.ObtenerRuidoBlanco)

	// Tipos de audio

	router.GET("/tipos-audio", audioController.ObtenerTiposDeAudio)


	router.Run(":8080")

}
