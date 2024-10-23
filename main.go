package main

import (
	"log"

	"gioui.org/app"
	"github.com/daanooo/request-capture/internal/server"
	"github.com/daanooo/request-capture/internal/ui"
)

func main() {
	captures := make(chan server.Capture) // Channel for sending captures to the UI

	l := server.NewListener(":55556", captures)
	ui := ui.NewUI(captures)

	// Start the application
	go l.Start()

	go func() {
		if err := ui.Start(); err != nil {
			log.Fatalf("Application terminated unexpectedly with error: %s\n", err)
		}
	}()

	app.Main()

	log.Fatalln("Application terminated normally")
}
