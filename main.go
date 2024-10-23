package main

import (
	"log"
	"os"

	"gioui.org/app"
	"github.com/daanooo/request-capture/internal/server"
	"github.com/daanooo/request-capture/internal/ui"
)

func main() {
	quit := make(chan error)              // Signalling channel to stop the application
	captures := make(chan server.Capture) // Channel for sending captures to the UI

	l := server.NewListener(":55556", quit, captures)
	ui := ui.NewUI(quit, captures)

	// Listen to the signalling channel and quit the application if the ui or listener requests it
	go func() {
		if err := <-quit; err != nil {
			log.Fatalf("Application shut down unexpectedly with error: %s\n", err)
		}

		log.Println("Application shut down gracefully")
		os.Exit(0)
	}()

	// Start the listener
	go l.Start()

	// Start the main application
	ui.Start()
	app.Main()
}
