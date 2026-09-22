package fachada

import (
	capaaccesodatos "almacenamiento/CapaAccesoADatos"
	componenteconexioncola "almacenamiento/CapaFachadaServices/ComponenteConexionCola"
	dtos "almacenamiento/CapaFachadaServices/DTOs"
	"fmt"
)

type FachadaAlmacenamiento struct {
	repo         *capaaccesodatos.RepositorioCanciones
	conexionCola *componenteconexioncola.RabbitPublisher
}

// Constructor de la fachada
func NuevaFachadaAlmacenamiento() *FachadaAlmacenamiento {
	fmt.Println(" Inicializando fachada de almacenamiento...")

	repo := capaaccesodatos.GetRepositorioCanciones()

	conexionCola, err := componenteconexioncola.NewRabbitPublisher()
	if err != nil {
		fmt.Println("Error al conectar con RabbitMQ:", err)
		conexionCola = nil
	}

	return &FachadaAlmacenamiento{
		repo:         repo,
		conexionCola: conexionCola,
	}
}

func (thisF *FachadaAlmacenamiento) GuardarCancion(objCancion dtos.CancionAlmacenamientoDTOInput, data []byte) error {
	thisF.conexionCola.PublicarNotificacion(componenteconexioncola.NotificacionCancion{
		Titulo:  objCancion.Titulo,
		Artista: objCancion.Artista,
		Genero:  objCancion.Genero,
		Mensaje: "Nueva canción almacenada: " + objCancion.Titulo + " de " + objCancion.Artista,
	})

	// Guardar archivo y registro en memoria
	// Delegar en el repositorio
	return thisF.repo.GuardarCancion(objCancion.Titulo, objCancion.Genero, objCancion.Artista, data)
}
