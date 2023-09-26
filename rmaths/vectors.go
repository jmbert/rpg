package rmaths

import (
	"math"
)

type Vec2 struct {
	X,
	Y float64
}

func (v *Vec2) Rotate(angle float64) Vec2 {
	var v2 Vec2

	v2.X = v.X*math.Cos(angle) - v.Y*math.Sin(angle)
	v2.Y = v.X*math.Sin(angle) + v.Y*math.Cos(angle)

	return v2
}
