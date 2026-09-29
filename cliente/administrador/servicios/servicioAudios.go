package servicios

import (
	"administrador/dtos"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// ServicioAudios consume el servicio REST del servidor de audios.
type ServicioAudios struct {
	urlBase    string
	clienteWeb *http.Client
}

// NuevoServicioAudios crea el servicio.
func NuevoServicioAudios(urlBase string) *ServicioAudios {
	return &ServicioAudios{urlBase: urlBase, clienteWeb: &http.Client{Timeout: 2 * time.Minute}}
}

// AlmacenarAudio envía el mp3, el tipo y los metadatos en un formulario
// multipart con POST /audios/almacenamiento.
func (this *ServicioAudios) AlmacenarAudio(rutaArchivo string, tipo string, metadatosDTO interface{}) (dtos.RespuestaAlmacenamientoDTO, error) {
	var respuestaDTO dtos.RespuestaAlmacenamientoDTO

	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		return respuestaDTO, fmt.Errorf("no se pudo abrir %s", rutaArchivo)
	}
	defer archivo.Close()

	metadatosJSON, err := json.Marshal(metadatosDTO)
	if err != nil {
		return respuestaDTO, err
	}

	// Armar el formulario: tipo, metadatos (JSON) y el archivo.
	var cuerpo bytes.Buffer
	formulario := multipart.NewWriter(&cuerpo)
	formulario.WriteField("tipo", tipo)
	formulario.WriteField("metadatos", string(metadatosJSON))
	campoArchivo, err := formulario.CreateFormFile("archivo", filepath.Base(rutaArchivo))
	if err != nil {
		return respuestaDTO, err
	}
	if _, err := io.Copy(campoArchivo, archivo); err != nil {
		return respuestaDTO, err
	}
	formulario.Close()

	respuesta, err := this.clienteWeb.Post(this.urlBase+"/audios/almacenamiento", formulario.FormDataContentType(), &cuerpo)
	if err != nil {
		return respuestaDTO, fmt.Errorf("no fue posible conectarse con el servidor de audios")
	}
	defer respuesta.Body.Close()

	json.NewDecoder(respuesta.Body).Decode(&respuestaDTO)
	if respuesta.StatusCode != http.StatusCreated {
		return respuestaDTO, fmt.Errorf("%s", respuestaDTO.Mensaje)
	}
	return respuestaDTO, nil
}
