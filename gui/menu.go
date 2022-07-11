package main

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type valButton struct {
	*widget.Clickable
	grid *Grid
}

type resetButton struct {
	*widget.Clickable
	grid *Grid
}

type solveButton struct {
	*widget.Clickable
	grid *Grid
}

func (g *Grid) newSolveButton() *solveButton {
	vb := new(solveButton)
	vb.grid = g
	vb.Clickable = new(widget.Clickable)
	return vb
}

func (g *Grid) newValButton() *valButton {
	vb := new(valButton)
	vb.grid = g
	vb.Clickable = new(widget.Clickable)
	return vb
}

func (g *Grid) newResetButton() *resetButton {
	sr := new(resetButton)
	sr.grid = g
	sr.Clickable = new(widget.Clickable)
	return sr
}

func (sr *resetButton) Layout(gtx layout.Context) layout.Dimensions {
	btn := material.Button(sr.grid.th, sr.Clickable, "Clear")
	if len(sr.Clicks()) != 0 {
		sr.grid.Reset()
	}
	btn.Background = menuColor
	return layout.UniformInset(10).Layout(gtx, btn.Layout)
}
func (sr *solveButton) Layout(gtx layout.Context) layout.Dimensions {
	btn := material.Button(sr.grid.th, sr.Clickable, "Solve")
	if len(sr.Clicks()) != 0 {
		sr.grid.Reset()
		sr.grid.current = sr.grid.solution.Clone()
	}
	btn.Background = menuColor
	return layout.UniformInset(10).Layout(gtx, btn.Layout)
}

func (vb *valButton) Layout(gtx layout.Context) layout.Dimensions {
	btn := material.Button(vb.grid.th, vb.Clickable, "Verify")
	if vb.Clickable.Pressed() {
		if vb.grid.current.Valid() {
			btn.Background = okColor
			btn.Text = "  OK !   "
		} else {
			btn.Background = notOkColor
			btn.Text = "Invalid !"
		}
	}
	btn.Background = menuColor
	return layout.UniformInset(10).Layout(gtx, btn.Layout)
}
