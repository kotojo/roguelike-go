package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/entity"
	"github.com/kotojo/roguelike_go/map_objects"
)

const BlockSize = 16
const FovRadius = 10
const MaxMonstersPerRoom = 3

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

	dejaVuFont := rl.LoadFont("DejaVuSansMono.ttf")

	player := &entity.Entity{
		X:      int(mapWidth) / 2,
		Y:      int(mapHeight) / 2,
		Char:   "@",
		Color:  rl.White,
		Name:   "Player",
		Blocks: true,
	}

	entities := []*entity.Entity{
		player,
	}

	gameMap := map_objects.NewGameMap(mapWidth, mapHeight)
	gameMap.MakeMap(maxRooms, roomMinSize, roomMaxSize, mapWidth, mapHeight, player, &entities, MaxMonstersPerRoom)

	fovRecompute := true

	gameState := PlayersTurn
	for !rl.WindowShouldClose() {
		// Movement
		action := handleKeys()

		fullscreen := action != nil && action.actionType == Fullscreen
		if fullscreen {
			rl.ToggleFullscreen()
		}

		movement := action != nil && action.actionType == Movement
		if movement && gameState == PlayersTurn {
			destinationX := player.X + action.actionMovement.dx
			destinationY := player.Y + action.actionMovement.dy
			isBlocked := gameMap.IsPlayerBlocked(destinationX, destinationY)

			if !isBlocked {
				target := entity.GetBlockingEntitiesAtLocation(entities, destinationX, destinationY)
				if target != nil {
					fmt.Printf("You kick %s in the shins, much to its annoyance \n", target.Name)
				} else {
					player.Move(
						action.actionMovement.dx,
						action.actionMovement.dy,
					)
					fovRecompute = true
				}
				gameState = EnemyTurn
			}
		}

		if fovRecompute {
			gameMap.RecomputeFov(player.X, player.Y, FovRadius)
		}

		if gameState == EnemyTurn {
			for _, entity := range entities {
				if entity != player {
					fmt.Printf("The %s ponders the meaning of it's existance", entity.Name)
				}
			}
			gameState = PlayersTurn
		}

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		renderAll(dejaVuFont, entities, gameMap)
		fovRecompute = false

		rl.EndDrawing()
	}

	rl.UnloadFont(dejaVuFont)

	rl.CloseWindow()
}
