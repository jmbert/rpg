package world

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"rpg/graphics"
	"strconv"
	"strings"
	"unicode"
)

var Actors map[string]Actor = make(map[string]Actor, 0)

type Assets struct {
	Map    MapJSON
	Actors []string
}

type MapJSON struct {
	Dat      string
	IndexMap map[string]string
}

type ActorJSON struct {
	Image string
}

func InterpretAssets() (World, error) {
	assetlocs, err := os.ReadFile("assets/assets.json")
	if err != nil {
		return World{}, err
	}

	var assets Assets

	err = json.Unmarshal(assetlocs, &assets)
	if err != nil {
		return World{}, err
	}
	var w World

	indexMap := make(map[int]Tile)

	for indexStr, tileName := range assets.Map.IndexMap {
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return w, err
		}
		tile, err := decipherTile("assets/tiles/" + tileName + ".json")
		if err != nil {
			return w, err
		}
		indexMap[index] = tile
	}

	var wmap Map
	DecipherWorldMap("assets/dat/maps/"+assets.Map.Dat+".map", &wmap, indexMap)
	w.Map = wmap

	for _, actorf := range assets.Actors {
		actordat, err := os.ReadFile("assets/actors/" + actorf + ".json")
		if err != nil {
			return w, err
		}

		var actorjson ActorJSON
		err = json.Unmarshal(actordat, &actorjson)
		if err != nil {
			return w, err
		}
		var actor Actor
		actor.image = getImage("assets/images/actors/" + actorjson.Image)

		Actors[actorf] = actor
	}

	return w, nil
}

type FlagsJSON struct {
	Walkable bool
}

func (f *FlagsJSON) toInt() TileFlags {
	var flags TileFlags

	if f.Walkable {
		flags += WALKABLE
	}

	return flags
}

type TileJSON struct {
	Image        string
	Flags        FlagsJSON
	MovementCost float64
}

func decipherTile(fileName string) (Tile, error) {

	file, err := os.ReadFile(fileName)
	if err != nil {
		return Tile{}, err
	}

	var tile TileJSON

	err = json.Unmarshal(file, &tile)
	if err != nil {
		return Tile{}, err
	}
	var t Tile

	t.Image = getImage("assets/images/tiles/" + tile.Image)

	t.movecost = tile.MovementCost

	t.Flags = tile.Flags.toInt()

	return t, nil
}

func DecipherWorldMap(file string, wmap *Map, indexMap map[int]Tile) {
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
		wmap.tiles = append(wmap.tiles, row)
	}

	for filecontents != "" {
		filecontents = decipherWS(filecontents)
		filecontents = decipherLine(filecontents, wmap, indexMap)
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

func decipherLine(line string, world *Map, indexMap map[int]Tile) string {
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

	world.tiles[x][y] = indexMap[index]
	return line[i:]
}

func getImage(file string) image.Image {
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
