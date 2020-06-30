package entities

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Entity struct {
	X           int
	Y           int
	Char        string
	Color       rl.Color
	Name        string
	Blocks      bool
	RenderOrder RenderOrder
	Fighter     *Fighter
	Ai          *BasicMonster
}

func NewPlayer(X, Y int) *Entity {
	playerFighter := &Fighter{
		Hp:      30,
		MaxHp:   30,
		Defense: 2,
		Power:   5,
	}
	return newEntity(X, Y, "@", rl.White, "Player", true, RenderOrderActor, playerFighter, nil)
}

func NewOrc(X, Y int) *Entity {
	monsterFighter := &Fighter{
		Hp:      10,
		MaxHp:   10,
		Defense: 0,
		Power:   3,
	}
	monsterAi := &BasicMonster{}
	return newEntity(X, Y, "O", rl.DarkGreen, "Orc", true, RenderOrderActor, monsterFighter, monsterAi)
}

func NewTroll(X, Y int) *Entity {
	monsterFighter := &Fighter{
		Hp:      16,
		MaxHp:   16,
		Defense: 1,
		Power:   4,
	}
	monsterAi := &BasicMonster{}
	return newEntity(X, Y, "T", rl.DarkGreen, "Troll", true, RenderOrderActor, monsterFighter, monsterAi)
}

func newEntity(X, Y int, Char string, Color rl.Color, Name string, Blocks bool, RenderOrder RenderOrder, Fighter *Fighter, Ai *BasicMonster) *Entity {
	e := Entity{X, Y, Char, Color, Name, Blocks, RenderOrder, nil, nil}
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
