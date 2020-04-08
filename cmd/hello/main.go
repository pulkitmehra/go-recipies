package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	grpc_hw "github.com/pulkitmehra/go-recipies/pkg/grpc/helloworld"
)

var (
	errorCh  = make(chan error, 1)
	signalCh = make(chan os.Signal, 1)
)

func main() {
	go grpc_hw.MainGrpc(errorCh)
	waitForEvent()
}

func waitForEvent() {

	signal.Notify(
		signalCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGQUIT,
	)
	fmt.Println("use ctr+c to quit")
	fmt.Println()
	select {
	case sig := <-signalCh:
		log.Printf("OS shutdown signal: %v", sig)
	case err := <-errorCh:
		log.Println(err)
		panic(err)
	}
	log.Println("Server Exit")
}
