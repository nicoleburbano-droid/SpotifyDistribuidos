package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"google.golang.org/grpc"
	capacontroladores "servidor.local/grpc-servidor/capaControladores"
	componenteconexioncola "servidor.local/grpc-servidor/capaFachadaServices/ComponenteConexionCola"
	"servidor.local/grpc-servidor/capalogger"
	"servidor.local/grpc-servidor/discovery"
	pb "servidor.local/grpc-servidor/serviciosCancion"
)

func main() {
	grpcPort := env("GRPC_PORT", "50051")
	listener, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("Error escuchando en el puerto: %v", err)
	}

	publisher, err := componenteconexioncola.NewRabbitPublisher()
	if err != nil {
		log.Printf("RabbitMQ no está disponible; el streaming continuará sin notificaciones: %v", err)
	} else {
		defer publisher.Cerrar()
	}

	grpcServer := grpc.NewServer()

	// Crear logger como singleton
	objLogger := capalogger.CrearUnicaInstanciaDelLogger()
	defer objLogger.CerrarArchivo()

	objControlador := capacontroladores.NewControladorServidor(objLogger, publisher)

	// Se registra el controlador que ofrece el procedimiento remoto
	pb.RegisterAudioServiceServer(grpcServer, objControlador)

	fmt.Printf("Servidor gRPC escuchando en :%s...\n", grpcPort)
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Error al iniciar servidor gRPC: %v", err)
		}
	}()
	advertisePort, err := strconv.Atoi(env("PROXY_PORT", "8080"))
	if err != nil || advertisePort < 1 || advertisePort > 65535 {
		log.Fatalf("PROXY_PORT inválido")
	}
	registro, err := discovery.Registrar(env("CONSUL_ADDR", "127.0.0.1:8500"), "servidorStreaming", env("SERVICE_HOST", "127.0.0.1"), advertisePort)
	if err != nil {
		log.Fatalf("no fue posible registrarse en Consul: %v", err)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	registro.Desregistrar()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stopped := make(chan struct{})
	go func() { grpcServer.GracefulStop(); close(stopped) }()
	select {
	case <-stopped:
	case <-ctx.Done():
		grpcServer.Stop()
	}
	_ = listener.Close()
}

func env(clave, porDefecto string) string {
	if valor := os.Getenv(clave); valor != "" {
		return valor
	}
	return porDefecto
}
