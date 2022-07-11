package main

import (
	"fmt"

	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/xavier268/gudoku/sdk"
)

type Grid struct {
	puzzle, current *sdk.Table
	th              *material.Theme
	lines           []*gridLine
}

func NewGrid(puzzle *sdk.Table) *Grid {
	g := new(Grid)
	g.current = puzzle.Clone()
	g.puzzle = puzzle.Clone()
	g.th = material.NewTheme(gofont.Collection())

	for li := 0; li < 9; li++ {
		line := g.addLine()
		for i := 0; i < 9; i++ {
			line.addElement(g)
		}
	}
	return g
}

func (g *Grid) Layout(gtx layout.Context) layout.Dimensions {
	flc := make([]layout.FlexChild, 0, 9)
	for _, li := range g.lines {
		liw := li.Layout
		flc = append(flc, layout.Rigid(liw))
	}
	return layout.Flex{Axis: layout.Vertical, Spacing: 0}.Layout(gtx, flc...)
}

func (gl *gridLine) Layout(gtx layout.Context) layout.Dimensions {
	flc := make([]layout.FlexChild, 0, 9)
	for _, ge := range gl.elts {
		gew := ge.Layout
		flc = append(flc, layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(1).Layout(gtx, gew)
			}))
	}
	return layout.Flex{Axis: layout.Horizontal, Spacing: 0}.Layout(gtx, flc...)
}

type gridLine struct {
	lpos int
	elts []*gridElement
}

func (g *Grid) addLine() *gridLine {
	gl := new(gridLine)
	gl.lpos = len(g.lines)
	gl.elts = make([]*gridElement, 0, 9)
	g.lines = append(g.lines, gl)
	return gl
}

type gridElement struct {
	pos int
	*Grid
	*widget.Clickable
}

func (gl *gridLine) addElement(g *Grid) {
	ge := new(gridElement)
	ge.Clickable = new(widget.Clickable)
	ge.pos = len(gl.elts)
	ge.Grid = g
	gl.elts = append(gl.elts, ge)
}

/*
func drawSquare(ops *op.Ops, color color.NRGBA, val int) layout.Dimensions {
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{}.Add(ops)

	return layout.Dimensions{Size: image.Pt(100, 100)}
}
*/

func (ge *gridElement) Layout(gtx layout.Context) layout.Dimensions {

	// Confine the area for pointer events.
	// defer clip.Rect(image.Rect(0, 0, 50, 50)).Push(gtx.Ops).Pop()
	/*
		pointer.InputOp{
			Tag:   b,
			Types: pointer.Press | pointer.Release | pointer.Enter | pointer.Leave,
		}.Add(gtx.Ops)
		area.Pop()
	*/

	btn := material.Button(ge.th, ge.Clickable, fmt.Sprint(ge.current.Get(ge.pos)))
	for _, c := range ge.Clicks() {

		v := ge.current.Get(ge.pos) // TODO - wrong calculation, use line !
		if c.Modifiers == 0 {
			v = (v + 1) % 10
		} else {
			v = (v + 9) % 10
		}
		fmt.Printf("#%d -> %d\n", ge.pos, v) // debug
		ge.current.Set(ge.pos, v)
	}

	return btn.Layout(gtx)
}
