package capacontroladores

import (
	dtos "almacenamiento/CapaFachadaServices/DTOs"
	capafachada "almacenamiento/CapaFachadaServices/Fachada"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

// tamanioMaximoArchivo es el tamaño máximo permitido para un mp3 (50 MB).
const tamanioMaximoArchivo = 50 << 20

// ControladorAlmacenamientoAudios atiende el servicio REST del administrador.
type ControladorAlmacenamientoAudios struct {
	fachada *capafachada.FachadaAlmacenamiento
}

// NuevoControladorAlmacenamientoAudios crea el controlador.
func NuevoControladorAlmacenamientoAudios(fachada *capafachada.FachadaAlmacenamiento) *ControladorAlmacenamientoAudios {
	return &ControladorAlmacenamientoAudios{fachada: fachada}
}

// AlmacenarAudio - POST /audios/almacenamiento (multipart/form-data)
// Campos: "archivo" (el .mp3), "tipo" (musica | audiolibro | podcast |
// ruidoBlanco) y "metadatos" (el DTO del tipo en JSON).
func (this *ControladorAlmacenamientoAudios) AlmacenarAudio(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n[REST] Invocado %s /audios/almacenamiento desde %s\n", r.Method, r.RemoteAddr)

	if r.Method != http.MethodPost {
		responder(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, tamanioMaximoArchivo)
	if err := r.ParseMultipartForm(tamanioMaximoArchivo); err != nil {
		responder(w, http.StatusBadRequest, "Formulario inválido o archivo muy grande")
		return
	}

	file, cabecera, err := r.FormFile("archivo")
	if err != nil {
		responder(w, http.StatusBadRequest, "Falta el campo \"archivo\"")
		return
	}
	defer file.Close()

	if strings.ToLower(filepath.Ext(cabecera.Filename)) != ".mp3" {
		responder(w, http.StatusBadRequest, "Solo se permiten archivos .mp3")
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		responder(w, http.StatusBadRequest, "No se pudo leer el archivo")
		return
	}

	tipo := r.FormValue("tipo")
	fmt.Printf("[REST] Tipo: %s, archivo recibido: %s (%d bytes)\n", tipo, cabecera.Filename, len(data))

	respuesta, err := this.fachada.AlmacenarAudio(tipo, []byte(r.FormValue("metadatos")), data)
	if err != nil {
		codigo := http.StatusBadGateway
		if errors.Is(err, capafachada.ErrDatosInvalidos) {
			codigo = http.StatusBadRequest
		} else if errors.Is(err, capafachada.ErrAudioExistente) {
			codigo = http.StatusConflict
		}
		responder(w, codigo, err.Error())
		return
	}

	escribirJSON(w, http.StatusCreated, respuesta)
}

// responder envía una respuesta con código y mensaje.
func responder(w http.ResponseWriter, codigo int, mensaje string) {
	escribirJSON(w, codigo, dtos.RespuestaAlmacenamientoDTO{Codigo: codigo, Mensaje: mensaje})
}

// escribirJSON escribe la respuesta e imprime el resultado en el servidor.
func escribirJSON(w http.ResponseWriter, codigo int, respuesta dtos.RespuestaAlmacenamientoDTO) {
	respuesta.Codigo = codigo
	fmt.Printf("[REST] Respuesta %d: %s\n", codigo, respuesta.Mensaje)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(codigo)
	json.NewEncoder(w).Encode(respuesta)
}
