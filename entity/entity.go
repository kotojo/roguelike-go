package entity

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ActionResultType int

const (
	Message ActionResultType = iota + 1
	Dead
)

type ActionResult struct {
	ResultType    ActionResultType
	ActionMessage string
	DeadEntity    *Entity
}

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

func (e *Entity) MoveTo(x, y int) {
	e.X = x
	e.Y = y
}

func (e *Entity) MoveAStar(gameMap Map, target *Entity) {
	l := AStarPath(gameMap, e.X, e.Y, target.X, target.Y)
	move := l.Front().Value.([]int)
	e.MoveTo(move[0], move[1])
}

func (e *Entity) DistanceTo(target *Entity) int {
	dx := target.X - e.X
	dy := target.Y - e.Y
	return int(math.Sqrt(float64(dx*dx + dy*dy)))
}

func GetBlockingEntitiesAtLocation(entities []*Entity, destinationX, destinationY int) *Entity {
	for _, entity := range entities {
		if entity.Blocks && entity.X == destinationX && entity.Y == destinationY {
			return entity
		}
	}
	return nil
}
