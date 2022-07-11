// Package main creates cross-platfom gui app for paying sudoku.
package main

import (
	"flag"
	"fmt"

	"math/rand"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/xavier268/gudoku/sdk"
)

var flagMaxDifficulty int
var flagVerbose bool

func init() {

	flag.IntVar(&flagMaxDifficulty, "difficulty", 9*9, "maximum allowed difficulty (number of blank values)")
	flag.IntVar(&flagMaxDifficulty, "d", 9*9, "shorthand for -difficulty")

	flag.BoolVar(&flagVerbose, "v", false, "print more detailed (verbose) information ")

}

func main() {
	flag.Parse()
	if flagVerbose {
		fmt.Println("Gudoku is a sudoku builder/solver - (c) Xavier Gandillot 2022")
	}
	go runMainWindow()
	app.Main()
}

// launch the main window
func runMainWindow() {

	// prepare puzzle and solution
	rand := rand.New(rand.NewSource(time.Now().UnixMicro()))
	puzzle, solution := sdk.BuildRandom(rand, 81-flagMaxDifficulty)
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
				fmt.Println(e.Err)
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
