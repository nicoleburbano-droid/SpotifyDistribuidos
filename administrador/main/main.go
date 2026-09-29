package main

import (
	"administrador/controladores"
	"administrador/servicios"
	"administrador/vistas"
	"os"
)

// Cliente del administrador: permite almacenar nuevos audios en el servidor
// de audios mediante REST.
func main() {
	urlAudios := os.Getenv("URL_AUDIOS")
	if urlAudios == "" {
		urlAudios = "http://localhost:5000"
	}

	servicio := servicios.NuevoServicioAudios(urlAudios)
	controlador := controladores.NuevoControladorAdministrador(servicio)
	vistas.MostrarMenuPrincipal(controlador)
}
