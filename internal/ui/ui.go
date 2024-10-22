package ui

import (
	"fmt"

	"gioui.org/app"
	"github.com/daanooo/request-capture/internal/server"
)

func Draw(captures chan server.Capture, closed chan error) {
	go func() {
		for {
			c := <-captures
			fmt.Println(c)
		}
	}()

	w := new(app.Window)
	w.Option(app.Title("Request Capture"))

	for {
		switch e := w.Event().(type) {

		case app.FrameEvent:

		case app.DestroyEvent:
			closed <- e.Err
		}
	}
}
