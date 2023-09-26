package main

import (
	"fmt"
	"rpg/graphics"
	"rpg/world"

	"github.com/veandco/go-sdl2/sdl"
)

var globalWorld world.World

var inputMap = graphics.KeyInput{
	"mbleft": move,
}

func move(event graphics.Event) {
	mbEvent, ok := event.(*sdl.MouseButtonEvent)
	if !ok {
		return
	}

	pCoordX, pCoordY := mbEvent.X, mbEvent.Y

	coordX, coordY := pCoordX/world.TileWidth, pCoordY/world.TileHeight

	dest := world.TileCoord{int(coordX), int(coordY)}

	globalWorld.GetPlayer().Dest = dest
}

func main() {
	graphics.Initialise(500, 500, 0, sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_ACCELERATED)

	globalWorld = world.NewWorld()

	world.DecipherWorldMap("assets/testmap.map", &globalWorld)

	player := world.NewActor(world.TileCoord{5, 0}, world.GetImage("assets/testplayer.png"))
	globalWorld.SetPlayer(player)

	enemy := world.NewActor(world.TileCoord{1, 2}, world.GetImage("assets/testenemy.png"))
	enemy.Dest = globalWorld.GetPlayer().GetPos()
	globalWorld.NewActor(enemy)
	tmp, _ := globalWorld.GetTile(world.TileCoord{2, 1})
	fmt.Printf("%v\n", tmp.Flags.Walkable())

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
