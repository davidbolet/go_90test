package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/davidbolet/go_90test/client-api/controller"
	"github.com/davidbolet/go_90test/client-api/parser"
	"github.com/davidbolet/go_90test/client-api/repository"
)

//Main start for application
func main() {
	fileToRead := "../ports.json"
	defaultServerPort := "8088"
	defaultGrpcAddr := ":9090"
	if len(os.Args) > 1 {
		fileToRead = os.Args[1]
	}
	if len(os.Args) > 2 {
		defaultGrpcAddr = os.Args[2]
	}
	repo := repository.NewPortRepository(defaultGrpcAddr)
	//Launch the parser goroutine
	go func() {
		parser := parser.NewPortParser(repo)
		parser.ReadAndParseFile(fileToRead)
	}()
	//Create the Rest Service
	portWebService := controller.NewPortWebService(repo)
	srv := &http.Server{
		Addr:    defaultServerPort,
		Handler: portWebService,
	}
	//Run it in its own goroutine
	go func() {
		if err := portWebService.Run(":" + defaultServerPort); err != nil {
			log.Fatalf("Error starting HTTP server: " + err.Error())
		}
	}()
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel // block until signal received
	log.Println("shutting down app because signal was received")
	srv.Shutdown(nil)

}
