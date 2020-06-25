package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func handleKeys() *Action {
	var action *Action

	if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyK) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 0,
			dy: -1,
		}}
	} else if rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressed(rl.KeyJ) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 0,
			dy: 1,
		}}
	} else if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressed(rl.KeyH) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: -1,
			dy: 0,
		}}
	} else if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressed(rl.KeyL) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 1,
			dy: 0,
		}}
	} else if rl.IsKeyPressed(rl.KeyY) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: -1,
			dy: -1,
		}}
	} else if rl.IsKeyPressed(rl.KeyU) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 1,
			dy: -1,
		}}
	} else if rl.IsKeyPressed(rl.KeyB) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: -1,
			dy: 1,
		}}
	} else if rl.IsKeyPressed(rl.KeyN) {
		action = &Action{actionType: Movement, actionMovement: &ActionMovement{
			dx: 1,
			dy: 1,
		}}
	} else if rl.IsKeyDown(rl.KeyLeftAlt) && rl.IsKeyPressed(rl.KeyEnter) {
		action = &Action{actionType: Fullscreen}
	}
	return action
}
