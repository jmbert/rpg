package graphics

import (
	"image/color"
)

func DrawLine(fx, fy, tx, ty int32, col color.Color) {

	r, g, b, a := col.RGBA()

	r *= 0xff
	g *= 0xff
	b *= 0xff
	a *= 0xff

	r /= 0xffff
	g /= 0xffff
	b /= 0xffff
	a /= 0xffff

	renderer.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
	renderer.DrawLine(fx, fy, tx, ty)
}
