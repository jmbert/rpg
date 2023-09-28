package world

import (
	"image"
	"image/color"
	"math"
	"rpg/graphics"
	"rpg/rmaths"
)

const MOVEFRAME = 7

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
	return math.Atan2(float64(t.Y), float64(t.X))
}

type Actor struct {
	image image.Image
}

type ActorInstance struct {
	Actor

	pos TileCoord

	path []TileInstance

	faceAngle float64

	Dest TileCoord
}

func (s *ActorInstance) Draw() {
	var startX = s.pos.X*TileWidth + TileWidth/2
	var startY = s.pos.Y*TileHeight + TileHeight/2

	var ActorWidth = float64(TileWidth) * 0.6
	var ActorHeight = float64(TileHeight) * 0.6

	imgWidth, imgHeight := s.image.Bounds().Max.X, s.image.Bounds().Max.Y

	for x := -ActorWidth / 2; x < ActorWidth/2; x += 0.5 {
		for y := -ActorHeight / 2; y < ActorHeight/2; y += 0.5 {
			vec := rmaths.Vec2{X: x, Y: y}
			rVec := vec.Rotate(s.faceAngle)
			graphics.SetPixel(int32(startX)+int32(math.Floor(float64(rVec.X))), int32(startY)+int32(math.Floor(float64(rVec.Y))), s.image.At(int(x+ActorWidth/2)*imgWidth/int(ActorWidth), int(y+ActorHeight/2)*imgHeight/int(ActorHeight)))
		}
	}
}

func (s *ActorInstance) Rotate(angle float64) {
	s.faceAngle = angle
}

func (s *ActorInstance) GetPos() TileCoord {
	return s.pos
}

func (s *ActorInstance) SetPos(pos TileCoord, world World) {
	if s.CheckPos(pos, world) {
		s.pos = pos
	}
}

func (s *ActorInstance) CheckPos(pos TileCoord, world World) bool {

	if pos.X < 0 || pos.Y < 0 {
		return false
	}

	posTile, err := world.GetTile(pos)
	if err != nil {
		return false
	}

	if !posTile.Flags.Walkable() {
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

func (s *ActorInstance) SetPath(a []TileInstance) {
	s.path = a
}

func (s *ActorInstance) GetPath() []TileInstance {
	return s.path
}

func NewActor(pos TileCoord, image image.Image) *ActorInstance {
	var actor = ActorInstance{pos: pos, Dest: pos}

	actor.image = image

	return &actor
}

func (a *ActorInstance) Update(w *World) {

	if Frame%MOVEFRAME == 0 {
		if w.GetPlayer() != a {
			a.Dest = w.GetPlayer().pos
		}
		a.MoveNext(w)
	}
}

func (a *ActorInstance) MoveNext(w *World) {
	path := w.FindPath(*a, a.Dest)

	if len(path) > 1 {
		direction := path[1].coords.Sub(a.pos)
		a.Rotate(direction.Angle())

		a.SetPos(path[1].coords, *w)
		a.SetPath(path[1:])
	}
}

func (a *ActorInstance) DrawPath() {
	lastTile := a.GetPos()
	for _, tile := range a.GetPath() {

		graphics.DrawLine(int32(lastTile.X*TileWidth+TileWidth/2), int32(lastTile.Y*TileHeight+TileHeight/2), int32(tile.coords.X*TileWidth+TileWidth/2), int32(tile.coords.Y*TileHeight+TileHeight/2), color.RGBA{255, 0, 0, 255})

		lastTile = tile.coords
	}
}
