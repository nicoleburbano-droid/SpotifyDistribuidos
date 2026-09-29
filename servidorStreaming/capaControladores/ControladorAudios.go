package capacontroladores

import (
	"fmt"
	capaFachada "servidor.local/grpc-servidor/capaFachada"
	componenteconexioncola "servidor.local/grpc-servidor/capaFachada/ComponenteConexionCola"
	pb "servidor.local/grpc-servidor/serviciosAudio"
)

type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
	Publisher *componenteconexioncola.RabbitPublisher
}

// Implementación del procedimiento remoto
func (s *ControladorServidor) AudioStream(req *pb.AudioRequest, stream pb.AudioService_AudioStreamServer) error {
	fmt.Printf("[gRPC] AudioStream: solicitud remota para reproducir %q\n", req.Filename)

	// Delegar la lógica de lectura y envío de chunks (fragmentos) a la fachada.
	return capaFachada.EnviarFragmentosAudio(req.Filename, stream, s.Publisher)

}
