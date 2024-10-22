package main

import (
	"log"

	"gioui.org/app"
	"github.com/daanooo/request-capture/internal/server"
	"github.com/daanooo/request-capture/internal/ui"
)

func main() {
	closed := make(chan error)            // Stop and error signaling channel
	captures := make(chan server.Capture) // Channel for sending captures to the UI

	l := server.NewListener(":55555", captures, closed)

	// Start the application
	go l.Start()
	go ui.Draw(captures, closed)
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
