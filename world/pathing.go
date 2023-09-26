package world

import (
	"math"

	"github.com/jbert/aoc/astar"
)

const TILE_UNWALKABLE = 100

type TileInstance struct {
	Tile

	world  *World
	coords TileCoord
}

func (t *TileInstance) Cost(w *World) float64 {
	for _, actor := range w.actors {
		if actor.GetPos() == t.coords {
			return TILE_UNWALKABLE
		}
	}
	return t.movecost
}

func (w *World) Neighbours(t TileInstance) []TileInstance {
	var neighbours []TileInstance

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			tile, err := w.GetTile(TileCoord{t.coords.X + x, t.coords.Y + y})
			if err == nil && tile.Flags.Walkable() {
				neighbours = append(neighbours, tile)
			}
		}
	}

	return neighbours
}

func (w *World) Weight(from TileInstance, to TileInstance) float64 {
	return math.Abs(float64(from.coords.X)-float64(to.coords.X)) + math.Abs(float64(from.coords.Y)-float64(to.coords.Y))
}

func (w *World) FindPath(a Actor, dest TileCoord) []TileInstance {
	apos, err := w.GetTile(a.GetPos())
	if err != nil {
		return []TileInstance{}
	}
	destpos, err := w.GetTile(dest)
	if err != nil {
		return []TileInstance{}
	}

	path, _ := astar.Astar[TileInstance](apos, destpos, w, func(v TileInstance) float64 { return v.Cost(w) })

	return path
}
