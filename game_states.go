package main

type GameState int

const (
	PlayersTurn GameState = iota
	EnemyTurn
)
