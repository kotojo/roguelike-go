package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func handleKeys() *Action {
	var action *Action

	if rl.IsKeyPressed(rl.KeyUp) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 0,
			dy: -1,
		}}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 0,
			dy: 1,
		}}
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: -1,
			dy: 0,
		}}
	} else if rl.IsKeyPressed(rl.KeyRight) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 1,
			dy: 0,
		}}
	} else if rl.IsKeyDown(rl.KeyLeftAlt) && rl.IsKeyPressed(rl.KeyEnter) {
		action = &Action{actionType: Fullscreen}
	}
	return action
}
