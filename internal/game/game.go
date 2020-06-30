package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/internal/game/entities"
	"github.com/kotojo/roguelike_go/internal/game/state"
)

const BlockSize = 16
const FovRadius = 10
const MaxMonstersPerRoom = 3

func StartGame() {
	var screenWidth int32 = 80 * BlockSize
	var screenHeight int32 = 50 * BlockSize
	var panelWidth int32 = screenWidth / 2
	var panelHeight int32 = 3 * BlockSize
	mapWidth := 80
	mapHeight := 43
	panelX := (screenWidth / 2) - (panelWidth / 2)
	panelY := int32(mapHeight) * BlockSize

	roomMaxSize := 10
	roomMinSize := 6
	maxRooms := 30

	rl.InitWindow(screenWidth, screenHeight, "Yet Another Roguelike Tutorial")

	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.SetTargetFPS(60)

	dejaVuFont := rl.LoadFont("DejaVuSansMono.ttf")

	player := entities.NewPlayer(int(mapWidth)/2, int(mapHeight)/2)

	ents := []*entities.Entity{
		player,
	}

	gameMap := NewGameMap(mapWidth, mapHeight)
	gameMap.MakeMap(maxRooms, roomMinSize, roomMaxSize, mapWidth, mapHeight, player, &ents, MaxMonstersPerRoom)

	fovRecompute := true

	gameState := state.PlayersTurn
	for !rl.WindowShouldClose() {
		// Movement
		action := handleKeys()

		fullscreen := action != nil && action.actionType == Fullscreen
		if fullscreen {
			rl.ToggleFullscreen()
		}

		var playerTurnResults []*state.ActionResult

		movement := action != nil && action.actionType == Movement
		if movement && gameState == state.PlayersTurn {
			destinationX := player.X + action.actionMovement.dx
			destinationY := player.Y + action.actionMovement.dy
			isBlocked := gameMap.IsPlayerBlocked(destinationX, destinationY)

			if !isBlocked {
				target := entities.GetBlockingEntitiesAtLocation(ents, destinationX, destinationY)
				if target != nil {
					attackResults := player.Fighter.Attack(target)
					playerTurnResults = append(playerTurnResults, attackResults...)
					for _, playerTurnResult := range playerTurnResults {
						resultType := playerTurnResult.ResultType
						if resultType == state.Message {
							fmt.Println(playerTurnResult.ActionMessage)
						}
						if resultType == state.Dead {
							var message string
							if player.Fighter.Hp <= 0 {
								message, gameState = entities.KillPlayer(player)
							} else {
								message = entities.KillMonster(target)
							}
							fmt.Println(message)
						}
					}
				} else {
					player.Move(
						action.actionMovement.dx,
						action.actionMovement.dy,
					)
					fovRecompute = true
				}
				if gameState != state.PlayerDead {
					gameState = state.EnemyTurn
				}
			}
		}

		if fovRecompute {
			gameMap.RecomputeFov(player.X, player.Y, FovRadius)
		}

		if gameState == state.EnemyTurn {
			for _, enemy := range ents {
				if enemy.Ai != nil {
					enemyTurnResults := enemy.Ai.TakeTurn(player, gameMap)
					for _, enemyTurnResult := range enemyTurnResults {
						resultType := enemyTurnResult.ResultType
						if resultType == state.Message {
							fmt.Println(enemyTurnResult.ActionMessage)
						}
						if resultType == state.Dead {
							var message string
							if player.Fighter.Hp <= 0 {
								message, gameState = entities.KillPlayer(player)
							} else {
								message = entities.KillMonster(enemy)
							}
							fmt.Println(message)
						}
					}
				}
			}
			if gameState != state.PlayerDead {
				gameState = state.PlayersTurn
			}
		}

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		renderAll(dejaVuFont, ents, player, gameMap, screenWidth, screenHeight, panelX, panelY, panelWidth, panelHeight)
		fovRecompute = false

		rl.EndDrawing()
	}

	rl.UnloadFont(dejaVuFont)

	rl.CloseWindow()
}
