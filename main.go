package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/entity"
	"github.com/kotojo/roguelike_go/map_objects"
	"github.com/kotojo/roguelike_go/render_order"
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

	playerFighter := &entity.Fighter{
		Hp:      30,
		Defense: 2,
		Power:   5,
	}
	player := entity.NewEntity(int(mapWidth)/2, int(mapHeight)/2, "@", rl.White, "Player", true, render_order.Actor, playerFighter, nil)

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

		var playerTurnResults []*entity.ActionResult

		movement := action != nil && action.actionType == Movement
		if movement && gameState == PlayersTurn {
			destinationX := player.X + action.actionMovement.dx
			destinationY := player.Y + action.actionMovement.dy
			isBlocked := gameMap.IsPlayerBlocked(destinationX, destinationY)

			if !isBlocked {
				target := entity.GetBlockingEntitiesAtLocation(entities, destinationX, destinationY)
				if target != nil {
					attackResults := player.Fighter.Attack(target)
					playerTurnResults = append(playerTurnResults, attackResults...)
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

		for _, playerTurnResult := range playerTurnResults {
			resultType := playerTurnResult.ResultType
			if resultType == entity.Message {
				fmt.Println(playerTurnResult.ActionMessage)
			}
			if resultType == entity.Dead {
				var message string
				if playerTurnResult.DeadEntity == player {
					message, gameState = killPlayer(player)
				} else {
					message = killMonster(playerTurnResult.DeadEntity)
				}
				fmt.Println(message)
			}
		}

		if gameState == EnemyTurn {
			for _, e := range entities {
				if e.Ai != nil {
					enemyTurnResults := e.Ai.TakeTurn(player, gameMap)
					for _, enemyTurnResult := range enemyTurnResults {
						resultType := enemyTurnResult.ResultType
						if resultType == entity.Message {
							fmt.Println(enemyTurnResult.ActionMessage)
						}
						if resultType == entity.Dead {
							var message string
							if enemyTurnResult.DeadEntity == player {
								message, gameState = killPlayer(player)
							} else {
								message = killMonster(enemyTurnResult.DeadEntity)
							}
							fmt.Println(message)
						}
					}
				}
			}
			if gameState != PlayerDead {
				gameState = PlayersTurn
			}
		}

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		renderAll(dejaVuFont, entities, player, gameMap, screenHeight)
		fovRecompute = false

		rl.EndDrawing()
	}

	rl.UnloadFont(dejaVuFont)

	rl.CloseWindow()
}
