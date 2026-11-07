package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	capaControladores "servidor.local/grpc-servidor/capaControladores"
	componenteconexioncola "servidor.local/grpc-servidor/capaFachada/ComponenteConexionCola"
	pb "servidor.local/grpc-servidor/serviciosAudio"
)

func main() {
	publisher, err := componenteconexioncola.NewRabbitPublisher()
	if err != nil {
		log.Printf("RabbitMQ no está disponible; el streaming continuará sin notificaciones: %v", err)
	} else {
		defer publisher.Cerrar()
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAudioServiceServer(grpcServer, &capaControladores.ControladorServidor{Publisher: publisher})

	fmt.Println("Servidor gRPC escuchando en :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
