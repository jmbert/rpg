package world

import (
	"fmt"
	"image"
	"math"
	"rpg/graphics"
	"rpg/rmaths"
)

const MOVEFRAME = 10

var ActorWidth = TileWidth * 0.6
var ActorHeight = TileHeight * 0.6

type TileCoord struct {
	X,
	Y int
}

func (t *TileCoord) Add(t2 TileCoord) TileCoord {
	return TileCoord{t.X + t2.X, t.Y + t2.Y}
}

func (t *TileCoord) Sub(t2 TileCoord) TileCoord {
	return TileCoord{t.X - t2.X, t.Y - t2.Y}
}

func (t *TileCoord) Angle() float64 {
	return math.Acos(float64(t.X) / math.Sqrt(float64(t.X*t.X)+float64(t.Y*t.Y)))
}

type Actor struct {
	pos TileCoord

	path []TileInstance

	image image.Image

	faceAngle float64

	Dest TileCoord
}

func (s *Actor) Draw() {
	var startX = s.pos.X*TileWidth + TileWidth/2
	var startY = s.pos.Y*TileHeight + TileHeight/2

	imgWidth, imgHeight := s.image.Bounds().Max.X, s.image.Bounds().Max.Y

	for x := -ActorWidth / 2; x < ActorWidth/2; x += 0.5 {
		for y := -ActorHeight / 2; y < ActorHeight/2; y += 0.5 {
			vec := rmaths.Vec2{X: x, Y: y}
			rVec := vec.Rotate(s.faceAngle)
			graphics.SetPixel(int32(startX)+int32(math.Floor(float64(rVec.X))), int32(startY)+int32(math.Floor(float64(rVec.Y))), s.image.At(int(x+ActorWidth/2)*imgWidth/int(ActorWidth), int(y+ActorHeight/2)*imgHeight/int(ActorHeight)))
		}
	}
}

func (s *Actor) Rotate(angle float64) {
	s.faceAngle = angle
}

func (s *Actor) GetPos() TileCoord {
	return s.pos
}

func (s *Actor) SetPos(pos TileCoord, world World) {
	if s.CheckPos(pos, world) {
		s.pos = pos
	}
}

func (s *Actor) CheckPos(pos TileCoord, world World) bool {

	if pos.X < 0 || pos.Y < 0 {
		return false
	}

	posTile, err := world.GetTile(pos)
	if err != nil {
		return false
	}

	if !posTile.Flags.Walkable() {
		fmt.Println("HELP")
		return false
	}

	for _, actor := range world.GetActors() {
		a := actor.GetPos()
		if a == pos {
			return false
		}
	}

	return true
}

func (s *Actor) SetPath(a []TileInstance) {
	s.path = a
}

func (s *Actor) GetPath() []TileInstance {
	return s.path
}

func NewActor(pos TileCoord, image image.Image) *Actor {
	return &Actor{pos: pos, image: image, Dest: pos}
}

func (a *Actor) Update(w *World) {

	if Frame%MOVEFRAME == 0 {
		a.MoveNext(w)
	}
}

func (a *Actor) MoveNext(w *World) {
	path := w.FindPath(*a, a.Dest)

	if len(path) > 1 {
		direction := path[1].coords.Sub(a.pos)
		a.Rotate(direction.Angle())

		a.SetPos(path[1].coords, *w)
		a.SetPath(path[1:])
	}
}
