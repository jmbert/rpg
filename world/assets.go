package world

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"rpg/graphics"
	"strconv"
	"strings"
	"unicode"
)

var TestTileWalkable = Tile{color.RGBA{R: 255, G: 255, B: 255, A: 255}, WALKABLE, 1}
var TestTileUnwalkable = Tile{color.RGBA{R: 127, G: 127, B: 127, A: 255}, 0, TILE_UNWALKABLE}

var indexMap = map[int]Tile{
	0: TestTileUnwalkable,
	1: TestTileWalkable,
}

func DecipherWorldMap(file string, world *World) {
	filecontentsb, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}
	filecontents := string(filecontentsb)

	def := filecontents[0:4]
	index, err := strconv.Atoi(def)
	if err != nil {
		log.Fatalln(err)
	}
	filecontents = filecontents[4:]

	TilesX = graphics.Width / TileWidth
	TilesY = graphics.Height / TileHeight

	for y := 0; y < TilesY; y++ {
		var row []Tile
		for x := 0; x < TilesX; x++ {
			row = append(row, indexMap[index])
		}
		world.tiles = append(world.tiles, row)
	}

	for filecontents != "" {
		filecontents = decipherWS(filecontents)
		filecontents = decipherLine(filecontents, world)
	}

	fmt.Printf("Deciphered world map %s\n", file)
}

func decipherWS(file string) string {
	i := 0
	for ; len(file) > i; i++ {
		if !unicode.IsSpace(rune(file[i])) {
			break
		}
	}

	return file[i:]
}

func decipherLine(line string, world *World) string {
	i := 0
	if line[i] != '[' {
		log.Fatalf("%d:Expected '[', found %c\n", i, line[i])
	}
	i++

	var coords string
	for ; line[i] != ']'; i++ {
		coords += string(line[i])
	}
	i++

	parts := strings.Split(coords, ":")
	if len(parts) != 2 {
		log.Fatalf("%d:Invalid coordinates %s\n", i, coords)
	}
	n1str := parts[0]
	n2str := parts[1]
	x, err := strconv.Atoi(n1str)
	if err != nil {
		log.Fatalln(err)
	}
	y, err := strconv.Atoi(n2str)
	if err != nil {
		log.Fatalln(err)
	}

	if line[i] != '-' {
		log.Fatalf("%d:Expected '-', found %c\n", i, line[i])
	}
	i++

	var indexStr string
	for ; len(line) > i; i++ {
		if unicode.IsSpace(rune(line[i])) {
			break
		}
		indexStr += string(line[i])
	}

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		log.Fatalln(err)
	}

	world.SetTile(indexMap[index], TileCoord{x, y})
	return line[i:]
}

func GetImage(file string) image.Image {
	var img image.Image
	f, err := os.Open(file)
	if err != nil {
		log.Println(err.Error())
		return image.Black
	}

	img, ffmt, err := image.Decode(f)
	if err != nil {
		log.Println(err.Error())
		return image.Black
	}

	fmt.Printf("Decoded image %s, format %s\n", file, ffmt)

	return img
}
