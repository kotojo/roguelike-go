package state

type GameState int

const (
	PlayersTurn GameState = iota
	EnemyTurn
	PlayerDead
)
