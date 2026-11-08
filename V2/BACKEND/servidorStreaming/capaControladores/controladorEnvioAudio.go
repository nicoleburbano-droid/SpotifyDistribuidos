package capacontroladores

import (
	"context"
	"google.golang.org/grpc/peer"
	"log"
	capafachadaservices "servidor.local/grpc-servidor/capaFachadaServices"
	componenteconexioncola "servidor.local/grpc-servidor/capaFachadaServices/ComponenteConexionCola"
	"servidor.local/grpc-servidor/capalogger"
	pb "servidor.local/grpc-servidor/serviciosCancion"
	"strings"
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
	clienteID := ObtenerIDCliente(req.GetClienteId(), stream.Context())
	//invocación a operación asincrona
	go thisC.logger.AlmacenarSolicitud(req.Titulo, direcionCliente)
	combined := strings.TrimSpace(req.GetTitulo())
	formato := strings.TrimPrefix(strings.ToLower(strings.TrimSpace(req.GetFormato())), ".")
	if formato != "" && !strings.HasSuffix(strings.ToLower(combined), "."+formato) {
		combined += "." + formato
	}
	if err := capafachadaservices.StreamAudioFile(
		combined,
		func(data []byte) error {
			return stream.Send(&pb.FragmentoCancion{Data: data})
		}); err != nil {
		return err
	}
	if thisC.publisher != nil {
		if err := thisC.publisher.PublicarNotificacion(componenteconexioncola.NotificacionCancion{ClienteID: clienteID, Titulo: req.GetTitulo(), Mensaje: "Reproducción completada desde " + direcionCliente}); err != nil {
			log.Printf("No se pudo registrar la reproducción de %q en RabbitMQ: %v", req.GetTitulo(), err)
		}
	}
	return nil

}

func ObtenerIDCliente(clienteID string, ctx context.Context) string {
	if clienteID = strings.TrimSpace(clienteID); clienteID != "" {
		return clienteID
	}
	return ObtenerDireccionCliente(ctx)
}

func ObtenerDireccionCliente(ctx context.Context) string {
	if p, ok := peer.FromContext(ctx); ok {
		return p.Addr.String()
	}
	return "desconocido"
}
