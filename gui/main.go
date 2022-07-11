// Package main creates cross-platfom gui app for paying sudoku.
package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/xavier268/gudoku/sdk"
)

func main() {
	go runMainWindow()
	app.Main()
}

// launch the main window
func runMainWindow() {
	w := app.NewWindow()
	var ops op.Ops
	g := Grid(*NewGrid(sdk.NewTable()))

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			if e.Err != nil {
				log.Fatal(e.Err)
			}
			os.Exit(0)
		case system.FrameEvent:
			// new frame context
			gtx := layout.NewContext(&ops, e)

			// draw to ops
			g.Layout(gtx)
			//layout.UniformInset(100).Layout(gtx, btn.Layout)

			// update display
			e.Frame(gtx.Ops)
		}
	}
}
