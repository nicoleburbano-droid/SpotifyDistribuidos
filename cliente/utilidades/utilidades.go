package utilidades

import (
	"fmt"
	"io"
	"log"
	"time"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"cliente.local/grpc-cliente/dtos"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	pb "servidorStreaming.local/grpc-servidor/serviciosAudio"
)

func DecodificarReproducir(reader io.Reader, canalSincronizacion chan struct{}) {
	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		log.Fatalf("error decodificando MP3: %v", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(canalSincronizacion)
	})))
}

func RecibirAudio(
	stream pb.AudioService_AudioStreamClient,
	writer *io.PipeWriter,
	canalSincronizacion chan struct{}) {
	noFragmento := 0
	for {
		fragmento, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Canción recibida completa.")
			writer.Close()
			break
		}
		if err != nil {
			log.Fatalf("Error recibiendo chunk: %v", err)
		}
		noFragmento++
		fmt.Printf("\n Fragmento #%d recibido (%d bytes) reproduciendo ...", noFragmento, len(fragmento.Data))

		if _, err := writer.Write(fragmento.Data); err != nil {
			log.Printf("Error escribiendo en pipe: %v", err)
			break
		}
	}
	// Esperar hasta que termine la reproducción
	<-canalSincronizacion
	fmt.Println("Reproducción finalizada.")
}

func SolicitarMetadata(tipo string, ctx context.Context) []dtos.MetadataAudioDTO {
	// Escapar el tipo para poder utilizarlo correctamente
	// como parámetro de la URL.
	tipoEscapado := url.PathEscape(tipo)

	serverURL := fmt.Sprintf(
		"http://localhost:8080/audios/tipo/%s",
		tipoEscapado,
	)

	// Crear cliente HTTP.
	httpClient := &http.Client{}

	// Crear la petición GET asociada al contexto recibido.
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		serverURL,
		nil,
	)
	if err != nil {
		fmt.Printf("Error al crear la petición HTTP: %v\n", err)
		return nil
	}

	// Realizar la petición REST.
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error al conectar con el servidor de metadatos: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	// El servidor indica que no se encontraron audios.
	if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf(
			"No se encontraron audios para el tipo '%s'.\n",
			tipo,
		)
		return nil
	}

	// Cualquier código diferente de 200 se considera
	// una respuesta no exitosa.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf(
			"Error al consultar los audios. Código HTTP: %d\n",
			resp.StatusCode,
		)
		return nil
	}

	// DTO que representa exactamente la respuesta del servidor REST.
	var respuesta dtos.RespuestaAudiosPorTipoDTO

	// Convertir el JSON recibido a una estructura Go.
	err = json.NewDecoder(resp.Body).Decode(&respuesta)
	if err != nil {
		fmt.Printf(
			"Error al interpretar la respuesta JSON: %v\n",
			err,
		)
		return nil
	}

	fmt.Printf("\n%s\n", respuesta.Mensaje)

	// Verificar si existen audios.
	if len(respuesta.VectorAudiosPorTipo) == 0 {
		fmt.Printf(
			"No hay audios disponibles para el tipo '%s'.\n",
			tipo,
		)
		return nil
	}

	return respuesta.VectorAudiosPorTipo
}