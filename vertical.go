package duit

import (
	"image"

	"9fans.net/go/draw"
)

type Vertical struct {
	Kids  []*Kid
	Split func(r image.Rectangle) (heights []int)

	size    image.Point
	heights []int
}

func (ui *Vertical) Layout(env *Env, r image.Rectangle, cur image.Point) image.Point {
	r.Min = image.Pt(0, cur.Y)
	heights := ui.Split(r)
	if len(heights) != len(ui.Kids) {
		panic("bad number of heights from split")
	}
	ui.heights = heights
	ui.size = image.ZP
	for i, k := range ui.Kids {
		p := image.Pt(0, ui.size.Y)
		size := k.UI.Layout(env, image.Rectangle{p, image.Pt(r.Dx(), heights[i])}, image.ZP)
		k.r = image.Rectangle{p, p.Add(size)}
		ui.size.Y += heights[i]
	}
	ui.size.X = r.Dx()
	return ui.size
}

func (ui *Vertical) Draw(env *Env, img *draw.Image, orig image.Point, m draw.Mouse) {
	kidsDraw(env, ui.Kids, ui.size, img, orig, m)
}

func (ui *Vertical) Mouse(env *Env, m draw.Mouse) (result Result) {
	return kidsMouse(env, ui.Kids, m)
}

func (ui *Vertical) Key(env *Env, orig image.Point, m draw.Mouse, k rune) (result Result) {
	return kidsKey(env, ui, ui.Kids, orig, m, k)
}

func (ui *Vertical) FirstFocus(env *Env) *image.Point {
	return kidsFirstFocus(env, ui.Kids)
}

func (ui *Vertical) Focus(env *Env, o UI) *image.Point {
	return kidsFocus(env, ui.Kids, o)
}

func (ui *Vertical) Print(indent int, r image.Rectangle) {
	uiPrint("Box", indent, r)
	kidsPrint(ui.Kids, indent+1)
}