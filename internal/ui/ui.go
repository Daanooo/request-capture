package ui

import (
	"fmt"
	"image/color"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/daanooo/request-capture/internal/server"
)

type UI struct {
	quit     chan error
	captures chan server.Capture
}

func NewUI(quit chan error, captures chan server.Capture) UI {
	return UI{quit, captures}
}

func (ui UI) Start() {
	window := new(app.Window)
	window.Option(app.Title("Request Capture"))
	window.Option(app.Size(unit.Dp(1280), unit.Dp(720)))

	ui.loop(window)
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

			fillBg(gtx.Ops)

			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			ui.quit <- e.Err
		}
	}
}

func fillBg(ops *op.Ops) {
	color := color.NRGBA{87, 87, 87, 255}
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
