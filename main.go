package main

import (
	"log"

	"github.com/daanooo/request-capture/internal/server"
)

func main() {
	l := server.NewListener(":55555")

	go l.Start()

	if err := <-l.Closed(); err != nil {
		log.Fatalf("Server unexpectedly closed with error: %s\n", err)
	}

	log.Println("Server quit gracefully")
}
