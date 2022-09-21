package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rombintu/timeshibot-backend/server"
)

func main() {
	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exitCh
		log.Println("Exit with 0")
		os.Exit(0)
	}()

	s := server.NewServer()
	if err := s.Start(); err != nil {
		log.Fatalf("%v+", err)
	}
}
