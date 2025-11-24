package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Paulinhoh/explicAi/configuration"
)

func main() {
	go configuration.NewAplication().Start()
	shutDown()
}

func shutDown() {
	signalShutDown := make(chan os.Signal, 2)
	signal.Notify(signalShutDown, syscall.SIGINT, syscall.SIGTERM)

	switch <-signalShutDown {
	case syscall.SIGINT:
		fmt.Println("SIGINT signal, explicAi is stoping...")
	case syscall.SIGTERM:
		fmt.Println("SIGTERM signal, explicAi is stoping...")
	}
}
