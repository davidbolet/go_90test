package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/davidbolet/go_90test/portdomsvc/pkg/repository"
	"github.com/davidbolet/go_90test/portdomsvc/pkg/service"
	"github.com/davidbolet/go_90test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Starting PortDomainService RPC server")
	defer log.Println("PortDomainService RPC server stopped")
	portListener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf(err.Error())
	}
	server := grpc.NewServer()
	reflection.Register(server)
	service := *service.CreateService(repository.CreateRepository())
	proto.RegisterPortDomainServiceServer(server, service)
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err = server.Serve(portListener); err != nil {
			log.Fatal(err.Error())
		}
	}()
	<-signalChannel // block until signal received
	log.Println("shutting down app because signal was received")
	server.GracefulStop()
	portListener.Close()

}
