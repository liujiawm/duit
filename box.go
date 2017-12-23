package duit

import (
	"image"

	"9fans.net/go/draw"
)

func NewBox(uis ...UI) *Box {
	kids := make([]*Kid, len(uis))
	for i, ui := range uis {
		kids[i] = &Kid{UI: ui}
	}
	return &Box{Kids: kids}
}

// box keeps elements on a line as long as they fit
type Box struct {
	Kids []*Kid

	size image.Point
}

func (ui *Box) Layout(env *Env, r image.Rectangle, ocur image.Point) image.Point {
	xmax := 0
	cur := image.Point{0, 0}
	nx := 0    // number on current line
	liney := 0 // max y of current line
	for _, k := range ui.Kids {
		p := k.UI.Layout(env, r, cur.Add(image.Pt(0, liney)))
		var kr image.Rectangle
		if nx == 0 || cur.X+p.X <= r.Dx() {
			kr = image.Rectangle{cur, cur.Add(p)}
			cur.X += p.X
			if p.Y > liney {
				liney = p.Y
			}
			nx += 1
		} else {
			cur.X = 0
			cur.Y += liney
			kr = image.Rectangle{cur, cur.Add(p)}
			nx = 1
			cur.X = p.X
			liney = p.Y
		}
		k.r = kr
		if xmax < cur.X {
			xmax = cur.X
		}
	}
	if len(ui.Kids) > 0 {
		cur.Y += liney
	}
	ui.size = image.Point{xmax, cur.Y}
	return ui.size
}

func (ui *Box) Draw(env *Env, img *draw.Image, orig image.Point, m draw.Mouse) {
	kidsDraw(env, ui.Kids, ui.size, img, orig, m)
}

func (ui *Box) Mouse(env *Env, m draw.Mouse) Result {
	return kidsMouse(env, ui.Kids, m)
}

func (ui *Box) Key(env *Env, orig image.Point, m draw.Mouse, c rune) Result {
	return kidsKey(env, ui, ui.Kids, orig, m, c)
}

func (ui *Box) FirstFocus(env *Env) *image.Point {
	return kidsFirstFocus(env, ui.Kids)
}

func (ui *Box) Focus(env *Env, o UI) *image.Point {
	return kidsFocus(env, ui.Kids, o)
}

func (ui *Box) Print(indent int, r image.Rectangle) {
	uiPrint("Box", indent, r)
	kidsPrint(ui.Kids, indent+1)
}