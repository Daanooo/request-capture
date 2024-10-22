package main

import (
	"log"

	"gioui.org/app"
	"github.com/daanooo/request-capture/internal/server"
	"github.com/daanooo/request-capture/internal/ui"
)

func main() {
	stop := make(chan error)              // Stop and error signaling channel
	captures := make(chan server.Capture) // Channel for sending captures to the UI

	l := server.NewListener(":55555", captures, stop)
	ui := ui.NewUI(stop, captures)

	// Start the application
	go l.Start()

	go func() {
		if err := ui.Start(); err != nil {
			stop <- err
		}
	}()

	go func() {
		app.Main()
		stop <- nil
	}()

	// Terminate the application when the UI or or listener signals or gives an error
	if err := <-stop; err != nil {
		log.Fatalf("Application terminated unexpectedly with error: %s\n", err)
	}

	log.Fatalln("Application terminated normally")
}
