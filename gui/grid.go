package main

import (
	"fmt"
	"image/color"

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
	g.puzzle = puzzle.Clone()
	g.Reset()
	return g
}

func (g *Grid) Reset() {
	g.current = g.puzzle.Clone()
	g.th = material.NewTheme(gofont.Collection())
	g.th.Palette.ContrastBg = color.NRGBA{
		R: 0x50,
		G: 0x90,
		B: 0x60,
		A: 0xff,
	}
	g.th.Palette.ContrastFg = color.NRGBA{
		R: 0,
		G: 0x30,
		B: 0,
		A: 0xff,
	}
	g.th.TextSize = 30

	for li := 0; li < 9; li++ {
		line := g.addLine()
		for i := 0; i < 9; i++ {
			line.addElement(g)
		}
	}
}

func (g *Grid) Valid() bool {
	return g.current.Valid()
}

func (g *Grid) Layout(gtx layout.Context) layout.Dimensions {
	flc := make([]layout.FlexChild, 0, 9)
	for _, li := range g.lines {
		liw := li.Layout
		flc = append(flc, layout.Flexed(1., liw))
	}
	return layout.Flex{Axis: layout.Vertical, Spacing: 0}.Layout(gtx, flc...)
}

func (gl *gridLine) Layout(gtx layout.Context) layout.Dimensions {
	flc := make([]layout.FlexChild, 0, 9)
	for _, ge := range gl.elts {
		gew := ge.Layout
		flc = append(flc, layout.Flexed(1.,
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
	ge.pos = gl.lpos*9 + len(gl.elts)
	ge.Grid = g
	gl.elts = append(gl.elts, ge)
}

func (ge *gridElement) Layout(gtx layout.Context) layout.Dimensions {

	for _, c := range ge.Clicks() {
		v := ge.current.Get(ge.pos)
		if c.Modifiers == 0 {
			v = (v + 1) % 10
		} else {
			v = (v + 9) % 10
		}
		fmt.Printf("#%d -> %d\n", ge.pos, v) // debug
		ge.current.Set(ge.pos, v)
	}
	if ge.current.Get(ge.pos) == 0 {
		btn := material.Button(ge.th, ge.Clickable, " ")
		btn.Background = color.NRGBA{
			R: 0x88,
			G: 0x55,
			B: 0x22,
			A: 0xff,
		}
		return btn.Layout(gtx)
	} else {
		btn := material.Button(ge.th, ge.Clickable, fmt.Sprint(ge.current.Get(ge.pos)))
		return btn.Layout(gtx)
	}
}
