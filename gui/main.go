// Package main creates cross-platfom gui app for paying sudoku.
package main

import (
	"log"
	"math/rand"
	"os"
	"time"

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

	// prepare puzzle and solution
	rand := rand.New(rand.NewSource(time.Now().UnixMicro()))
	puzzle, solution := sdk.BuildRandom(rand, 0)
	s := sdk.NewShuffler(rand)
	s.Shuffle(puzzle, solution)

	// create gui
	w := app.NewWindow()
	var ops op.Ops
	g := Grid(*NewGrid(puzzle, solution))
	vb := g.newValButton()
	sr := g.newResetButton()
	sv := g.newSolveButton()

	// main event loop
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
			layout.Flex{Axis: layout.Vertical}.Layout(
				gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
							layout.Rigid(sr.Layout),
							layout.Rigid(sv.Layout),
							layout.Rigid(vb.Layout))
					}),
				layout.Rigid(g.Layout),
			)

			// update display
			e.Frame(gtx.Ops)
		}
	}
}
