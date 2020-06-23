package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	X      int
	Y      int
	Char   string
	Color  rl.Color
	Name   string
	Blocks bool
}

func (e *Entity) Move(x, y int) {
	e.X += x
	e.Y += y
}

func GetBlockingEntitiesAtLocation(entities []*Entity, destinationX, destinationY int) *Entity {
	for _, entity := range entities {
		if entity.Blocks && entity.X == destinationX && entity.Y == destinationY {
			return entity
		}
	}
	return nil
}
