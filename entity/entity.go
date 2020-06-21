package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	X     int
	Y     int
	Char  string
	Color rl.Color
}

func (e *Entity) Move(x, y int) {
	e.X += x
	e.Y += y
}
