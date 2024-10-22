package main

import (
	"log"

	"gioui.org/app"
	"github.com/daanooo/request-capture/internal/server"
	"github.com/daanooo/request-capture/internal/ui"
)

func main() {
	// This channel is used to signal that the application should stop, either by the UI or the listener
	closed := make(chan error)
	l := server.NewListener(":55555", closed)

	// Start the application
	go l.Start()
	go ui.Draw(closed)
	go func() {
		app.Main()
		closed <- nil
	}()

	// Terminate the application when the UI or or listener signals or gives an error
	if err := <-closed; err != nil {
		log.Fatalf("Application terminated unexpectedly with error: %s\n", err)
	}

	log.Fatalln("Application terminated normally")
}
