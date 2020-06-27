package game

type GameState int

const (
	PlayersTurn GameState = iota
	EnemyTurn
	PlayerDead
)
