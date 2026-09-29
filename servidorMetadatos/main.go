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

	router.POST("/audios", audioController.RegistrarAudio)
	router.GET("/audios/titulo/:titulo", audioController.ConsultarAudio)
	router.GET("audios/tipo/:tipo", audioController.ConsultarTodos)

	router.Run(":8080")

}
