package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/entity"
	"github.com/kotojo/roguelike_go/map_objects"
)

const BlockSize = 16
const FovRadius = 10

func main() {
	var screenWidth int32 = 80 * BlockSize
	var screenHeight int32 = 50 * BlockSize
	mapWidth := 80
	mapHeight := 45

	roomMaxSize := 10
	roomMinSize := 6
	maxRooms := 30

	rl.InitWindow(screenWidth, screenHeight, "Yet Another Roguelike Tutorial")

	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.SetTargetFPS(60)

	pixelFont := rl.LoadFont("pixelplay.png")

	player := &entity.Entity{
		X:     int(mapWidth) / 2,
		Y:     int(mapHeight) / 2,
		Char:  "@",
		Color: rl.White,
	}

	npc := &entity.Entity{
		X:     int(mapWidth)/2 - 3,
		Y:     int(mapHeight) / 2,
		Char:  "@",
		Color: rl.Brown,
	}

	entities := []*entity.Entity{
		player,
		npc,
	}

	gameMap := map_objects.NewGameMap(mapWidth, mapHeight)
	gameMap.MakeMap(maxRooms, roomMinSize, roomMaxSize, mapWidth, mapHeight, player)

	canFullscreen := true
	canMove := true
	fovRecompute := true

	for !rl.WindowShouldClose() {
		// Movement
		action := handleKeys()

		// keypress gets counted multiple times no matter how quick you press the button.
		// This makes it so we should only full screen one time before waiting for a different
		// action to fire to allow fullscreening again.
		// Otherwise we end up calling `ToggleFullscreen` half a dozen times in < 1 sec
		if action == nil {
			canFullscreen = true
			canMove = true
		} else if action.actionType == Fullscreen && canFullscreen {
			rl.ToggleFullscreen()
			canFullscreen = false
		} else if action.actionType == Movement && canMove {
			isBlocked := gameMap.IsPlayerBlocked(player.X+action.actionMovement.dx, player.Y+action.actionMovement.dy)

			if !isBlocked {
				player.Move(
					action.actionMovement.dx,
					action.actionMovement.dy,
				)
			}
			canMove = false
			fovRecompute = true
		}

		if fovRecompute {
			gameMap.RecomputeFov(player.X, player.Y, FovRadius)
		}

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		renderAll(pixelFont, entities, gameMap)
		fovRecompute = false

		rl.EndDrawing()
	}

	rl.UnloadFont(pixelFont)

	rl.CloseWindow()
}
