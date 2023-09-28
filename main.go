package main

import (
	"log"
	"os"
	"rpg/graphics"
	"rpg/world"

	"github.com/veandco/go-sdl2/sdl"
)

var globalWorld *world.World

var inputMap = graphics.KeyInput{
	"mbleft": move,
}

func move(event graphics.Event) {
	mbEvent, ok := event.(*sdl.MouseButtonEvent)
	if !ok {
		return
	}

	pCoordX, pCoordY := mbEvent.X, mbEvent.Y

	coordX, coordY := int(pCoordX)/world.TileWidth, int(pCoordY)/world.TileHeight

	dest := world.TileCoord{X: int(coordX), Y: int(coordY)}

	globalWorld.GetPlayer().Dest = dest
}

func main() {
	graphics.Initialise(750, 750, 0, sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_ACCELERATED, os.Args[0])

	wworld, err := world.InterpretAssets()
	globalWorld = &wworld

	if err != nil {
		log.Fatalln(err)
	}

	globalWorld.NewActor(world.Actors["testplayer"], world.TileCoord{X: 1, Y: 0})

	globalWorld.NewActor(world.Actors["testenemy"], world.TileCoord{X: 1, Y: 9})

	for {
		code := graphics.MainloopIter(inputMap)
		if code == 1 {
			break
		}

		globalWorld.Update()
		globalWorld.Draw()
	}

	graphics.Destroy()
}
