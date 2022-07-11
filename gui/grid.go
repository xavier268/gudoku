package main

import (
	"fmt"

	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/xavier268/gudoku/sdk"
)

type Grid struct {
	puzzle, current, solution *sdk.Table
	th                        *material.Theme
	lines                     []*gridLine
}

func NewGrid(puzzle, solution *sdk.Table) *Grid {
	g := new(Grid)
	g.puzzle = puzzle.Clone()
	g.solution = solution.Clone()
	g.Reset()
	return g
}

func (g *Grid) Reset() {
	g.current = g.puzzle.Clone()
	g.th = material.NewTheme(gofont.Collection())
	g.th.Palette.ContrastBg = contrastBG
	g.th.Palette.ContrastFg = contratFG
	g.th.TextSize = unit.Sp(flagFontSize)
	g.lines = nil
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
	for i, li := range g.lines {
		if i == 3 || i == 6 {
			flc = append(flc, layout.Flexed(0.2, layout.Spacer{Width: 4, Height: 4}.Layout))
		}
		liw := li.Layout
		flc = append(flc, layout.Flexed(1., liw))
	}
	return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceSides}.Layout(gtx, flc...)
}

func (gl *gridLine) Layout(gtx layout.Context) layout.Dimensions {
	flc := make([]layout.FlexChild, 0, 9)
	for i, ge := range gl.elts {
		if i == 3 || i == 6 {
			flc = append(flc, layout.Flexed(0.2, layout.Spacer{Width: 4, Height: 4}.Layout))
		}
		gew := ge.Layout
		flc = append(flc, layout.Flexed(1.,
			func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(1).Layout(gtx, gew)
			}))
	}
	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceSides}.Layout(gtx, flc...)
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

	if ge.puzzle.Get(ge.pos) == 0 { // only if value can be modified, do not modify initial puzzle values.
		for _, c := range ge.Clicks() {
			v := ge.current.Get(ge.pos)
			if c.Modifiers == 0 {
				v = (v + 1) % 10
			} else {
				v = (v + 9) % 10
			}
			if flagVerbose {
				fmt.Printf("#%d -> %d\n", ge.pos, v)
			}
			ge.current.Set(ge.pos, v)
		}
	}

	btn := material.Button(ge.th, ge.Clickable, "")
	switch {

	case ge.puzzle.Get(ge.pos) != 0:
		btn = material.Button(ge.th, ge.Clickable, fmt.Sprint(ge.current.Get(ge.pos)))
		btn.Background = specialBG
	case ge.puzzle.Get(ge.pos) == 0 && ge.current.Get(ge.pos) != 0:
		btn = material.Button(ge.th, ge.Clickable, fmt.Sprint(ge.current.Get(ge.pos)))
	}

	return btn.Layout(gtx)
}
