package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func handleKeys() *Action {
	var action *Action

	if rl.IsKeyDown(rl.KeyUp) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 0,
			dy: -1,
		}}
	} else if rl.IsKeyDown(rl.KeyDown) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 0,
			dy: 1,
		}}
	} else if rl.IsKeyDown(rl.KeyLeft) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: -1,
			dy: 0,
		}}
	} else if rl.IsKeyDown(rl.KeyRight) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 1,
			dy: 0,
		}}
	} else if rl.IsKeyDown(rl.KeyEnter) && rl.IsKeyDown(rl.KeyLeftAlt) {
		action = &Action{actionType: Fullscreen}
	}
	return action
}
