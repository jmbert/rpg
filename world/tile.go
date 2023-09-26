package world

import (
	"image/color"
	"rpg/graphics"
)

type TileFlags int32

func (f TileFlags) Walkable() bool {
	return f&WALKABLE != 0
}

func (f TileFlags) PlayerTarget() bool {
	return f&PLAYERTARGET != 0
}

func (f TileFlags) Flags(flag TileFlags) {
	f |= flag
}

const (
	WALKABLE     = 1 << 0
	PLAYERTARGET = 1 << 1
)

const TileWidth = 50
const TileHeight = 50

var BorderColour = color.RGBA{0, 0, 0, 255}

type Tile struct {
	Colour color.RGBA

	Flags TileFlags

	movecost float64
}

func (t *Tile) Draw(xc, yc int, flags TileFlags) {
	startX := xc * TileWidth
	startY := yc * TileHeight

	allFlags := t.Flags | flags

	for x := 0; x < TileWidth; x++ {
		for y := 0; y < TileHeight; y++ {
			graphics.SetPixel(int32(x+startX), int32(y+startY), t.Colour)

			if x > TileWidth*39/40 || y > TileHeight*39/40 || x < TileWidth*1/40 || y < TileHeight*1/40 {
				if allFlags.PlayerTarget() {
					graphics.SetPixel(int32(x+startX), int32(y+startY), color.RGBA{255, 0, 0, 255})
				} else {
					graphics.SetPixel(int32(x+startX), int32(y+startY), BorderColour)
				}
			}
		}
	}
}
