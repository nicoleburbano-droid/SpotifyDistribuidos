package componenteclientemetadatos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ClienteMetadatos consume los servicios REST del servidor de metadatos para
// registrar los metadatos de un audio nuevo.
type ClienteMetadatos struct {
	urlBase    string
	clienteWeb *http.Client
}

// respuestaRegistro es la parte de la respuesta del servidor de metadatos que
// se necesita (tiene los campos "codigo" y "mensaje").
type respuestaRegistro struct {
	Codigo  int    `json:"codigo"`
	Mensaje string `json:"mensaje"`
}

// NuevoClienteMetadatos crea el cliente REST.
func NuevoClienteMetadatos(urlBase string) *ClienteMetadatos {
	return &ClienteMetadatos{urlBase: urlBase, clienteWeb: &http.Client{Timeout: 10 * time.Second}}
}

// RegistrarMetadatos hace POST al endpoint del tipo (/musica, /audiolibros,
// /podcasts o /ruido-blanco) enviando el DTO en JSON.
func (this *ClienteMetadatos) RegistrarMetadatos(endpoint string, metadatosDTO interface{}) error {
	cuerpo, err := json.Marshal(metadatosDTO)
	if err != nil {
		return fmt.Errorf("error convirtiendo los metadatos a JSON: %v", err)
	}

	url := this.urlBase + endpoint
	fmt.Printf("[REST] Invocando POST %s\n", url)

	respuesta, err := this.clienteWeb.Post(url, "application/json", bytes.NewReader(cuerpo))
	if err != nil {
		return fmt.Errorf("el servidor de metadatos no está disponible")
	}
	defer respuesta.Body.Close()

	var objRespuesta respuestaRegistro
	json.NewDecoder(respuesta.Body).Decode(&objRespuesta)

	if respuesta.StatusCode != http.StatusCreated {
		return fmt.Errorf("el servidor de metadatos respondió %d: %s", respuesta.StatusCode, objRespuesta.Mensaje)
	}
	fmt.Printf("[REST] Respuesta del servidor de metadatos: %s\n", objRespuesta.Mensaje)
	return nil
}
