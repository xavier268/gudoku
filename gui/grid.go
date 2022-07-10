package main

import (
	"fmt"
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type GridElement struct {
	*widget.Clickable
	val int
}

func NewGridElement() *GridElement {
	g := new(GridElement)
	g.Clickable = new(widget.Clickable)
	return g
}

/*
func drawSquare(ops *op.Ops, color color.NRGBA, val int) layout.Dimensions {
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	paint.ColorOp{Color: color}.Add(ops)
	paint.PaintOp{}.Add(ops)

	return layout.Dimensions{Size: image.Pt(100, 100)}
}
*/

func (b *GridElement) Layout(gtx layout.Context) layout.Dimensions {

	// Confine the area for pointer events.
	defer clip.Rect(image.Rect(0, 0, 50, 50)).Push(gtx.Ops).Pop()
	/*
		pointer.InputOp{
			Tag:   b,
			Types: pointer.Press | pointer.Release | pointer.Enter | pointer.Leave,
		}.Add(gtx.Ops)
		area.Pop()
	*/

	btn := material.Button(TH, b.Clickable, fmt.Sprint(b.val))
	for _, c := range b.Clicks() {

		b.val = (b.val + 1) % 10
		fmt.Printf("%x -> %d\n", c.Modifiers, b.val)
	}

	return btn.Layout(gtx)
}
