package ui

import (
	"gioui.org/app"
)

func Draw(closed chan error) {
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
