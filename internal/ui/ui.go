package ui

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/daanooo/request-capture/internal/server"
)

type UI struct {
	captures chan server.Capture
}

func NewUI(captures chan server.Capture) UI {
	return UI{captures}
}

func (ui UI) Start() error {
	window := new(app.Window)
	window.Option(app.Title("Request Capture"))
	window.Option(app.Size(unit.Dp(1280), unit.Dp(720)))

	if err := ui.loop(window); err != nil {
		return err
	}

	return nil
}

func (ui UI) loop(window *app.Window) error {
	go func() {
		for {
			c := <-ui.captures
			fmt.Println(c)
		}
	}()

	ops := new(op.Ops)

	for {
		switch e := window.Event().(type) {

		case app.FrameEvent:
			gtx := app.NewContext(ops, e)

			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}
