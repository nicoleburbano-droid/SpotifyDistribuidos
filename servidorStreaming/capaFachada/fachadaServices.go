package fachada

import (
	"fmt"
	"io"
	"log"
	"os"

	capaaccesodatos "servidor.local/grpc-servidor/capaAccesoDatos"
	pb "servidor.local/grpc-servidor/serviciosAudio"
)

// abrirArchivo actua como fachada para obtener el archivo de la cancion usando la capa de acceso a datos.
// Recibe el titulo (ruta) y devuelve el *os.File abierto o un error.
func abrirArchivo(titulo string) (*os.File, error) {
	log.Printf("GetAudioFile llamado con titulo=%s", titulo)
	return capaaccesodatos.AbrirArchivo(titulo)
}

// EnviarFragmentosAudio lee el archivo de la canción en chunks y los envía al stream gRPC.
// Esta función encapsula la lógica de lectura y envío para mantener el servidor limpio.
func EnviarFragmentosAudio(titulo string, stream pb.AudioService_AudioStreamServer) error {
	log.Printf("Enviando fragmentos de audio para titulo=%s", titulo)
	file, err := abrirArchivo(titulo)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	buf := make([]byte, 32*1024) // 32KB por chunk (fragmento)
	chunkNum := 0

	for {
		//n es la cantidad de bytes que devolvió file.Read(buf).
		// en buf se guarda el fragmento leido
		n, err := file.Read(buf)
		if err == io.EOF {
			log.Println("Cancion enviada completa.")
			break
		}
		if err != nil {
			return fmt.Errorf("error leyendo archivo: %w", err)
		}

		chunkNum++

		if n > 0 {
			// se crea un objeto de tipo AudioChunk
			objChunk := &pb.AudioChunk{Data: buf[:n]}
			//con la funcion send se envia el objeto de tipo AudioChunk
			if err := stream.Send(objChunk); err != nil {
				return fmt.Errorf("Error enviando chunk #%d: %w", chunkNum, err)
			}
			log.Printf("Chunk #%d enviado (%d bytes)\n", chunkNum, n)
		}
	}

	return nil
}
