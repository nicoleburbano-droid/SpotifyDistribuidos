package capacontroladores

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
	"log"
	capafachadaservices "servidor.local/grpc-servidor/capaFachadaServices"
	componenteconexioncola "servidor.local/grpc-servidor/capaFachadaServices/ComponenteConexionCola"
	"servidor.local/grpc-servidor/capalogger"
	pb "servidor.local/grpc-servidor/serviciosCancion"
)

type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
	logger    *capalogger.Logger
	publisher *componenteconexioncola.RabbitPublisher
}

func NewControladorServidor(logger *capalogger.Logger, publisher *componenteconexioncola.RabbitPublisher) *ControladorServidor {
	return &ControladorServidor{
		logger:    logger,
		publisher: publisher,
	}
}

// Implementación del procedimiento remoto
func (thisC *ControladorServidor) EnviarCancionMedianteStream(
	req *pb.PeticionDTO, stream pb.AudioService_EnviarCancionMedianteStreamServer) error {

	//invocación a operación sincrona
	direcionCliente := ObtenerDireccionCliente(stream.Context())
	//invocación a operación asincrona
	go thisC.logger.AlmacenarSolicitud(req.Titulo, direcionCliente)
	if thisC.publisher != nil {
		if err := thisC.publisher.PublicarNotificacion(componenteconexioncola.NotificacionCancion{Titulo: req.GetTitulo(), Mensaje: "Reproducción solicitada desde " + direcionCliente}); err != nil {
			log.Printf("No se pudo notificar la reproducción de %q a RabbitMQ: %v", req.GetTitulo(), err)
		}
	}
	combined := req.GetTitulo() + "." + req.GetFormato()
	return capafachadaservices.StreamAudioFile(
		combined,
		func(data []byte) error {
			return stream.Send(&pb.FragmentoCancion{Data: data})
		})

}

func ObtenerDireccionCliente(ctx context.Context) string {
	if p, ok := peer.FromContext(ctx); ok {
		return p.Addr.String()
	}
	return "desconocido"
}
