package main

import "image/color"

var contratFG = color.NRGBA{ // foreground text in grid
	R: 0,
	G: 0x30,
	B: 0,
	A: 0xff,
}

var contrastBG = color.NRGBA{ // background  for editable grid element
	R: 0x50,
	G: 0x90,
	B: 0x60,
	A: 0xff,
}

var specialBG = color.NRGBA{ // background for fixed grid element
	R: 0x88,
	G: 0x55,
	B: 0x22,
	A: 0xff,
}

var okColor = color.NRGBA{
	R: 0,
	G: 0x88,
	B: 0,
	A: 0xff,
}

var notOkColor = color.NRGBA{
	R: 0x88,
	G: 0,
	B: 0,
	A: 0xff,
}

var menuColor = color.NRGBA{
	R: 0x88,
	G: 0x88,
	B: 0x88,
	A: 0xff,
}
