package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Entity struct {
	X       int
	Y       int
	Char    string
	Color   rl.Color
	Name    string
	Blocks  bool
	Fighter *Fighter
	Ai      *BasicMonster
}

func NewEntity(X, Y int, Char string, Color rl.Color, Name string, Blocks bool, Fighter *Fighter, Ai *BasicMonster) *Entity {
	e := Entity{X, Y, Char, Color, Name, Blocks, nil, nil}
	if Fighter != nil {
		e.Fighter = Fighter
		e.Fighter.Owner = &e
	}
	if Ai != nil {
		e.Ai = Ai
		e.Ai.Owner = &e
	}
	return &e
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
