package world

import (
	"errors"
)

var TilesX int
var TilesY int

var Frame int64

type World struct {
	tiles  [][]Tile
	actors []*Actor

	player int
}

func NewWorld() World {
	var w World

	return w
}

func (w *World) Update() {
	Frame++
	for _, actor := range w.actors {
		actor.Update(w)
	}

}

func (w *World) Draw() {
	for x, row := range w.tiles {
		for y, tile := range row {
			var flags TileFlags
			coord := TileCoord{x, y}
			if coord == w.GetPlayer().Dest {
				flags = PLAYERTARGET
			}
			tile.Draw(x, y, flags)
		}
	}
	for _, actor := range w.actors {
		actor.Draw()
	}
}

func (w *World) SetTile(t Tile, coords TileCoord) {
	w.tiles[coords.X][coords.Y] = t
}

func (w *World) SetTileFlags(flags TileFlags, coords TileCoord) {
	w.tiles[coords.X][coords.Y].Flags = flags
}

func (w *World) NewActor(a *Actor) {
	w.actors = append(w.actors, a)
}

func (w *World) SetPlayer(p *Actor) {
	w.player = len(w.actors)
	w.NewActor(p)
}

func (w *World) GetPlayer() *Actor {
	return w.GetActors()[w.player]
}

func (w *World) GetActors() []*Actor {
	return w.actors
}

func (w *World) GetTile(pos TileCoord) (TileInstance, error) {
	if pos.X >= len(w.tiles) {
		return TileInstance{}, errors.New("out of bounds")
	}
	if pos.X < 0 || pos.Y < 0 || pos.Y >= len(w.tiles[pos.X]) {
		return TileInstance{}, errors.New("out of bounds")
	}
	return TileInstance{w.tiles[pos.X][pos.Y], w, pos}, nil
}