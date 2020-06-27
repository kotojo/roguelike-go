package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	playerFighter := &Fighter{
		Hp:      30,
		MaxHp:   30,
		Defense: 2,
		Power:   5,
	}
	player := NewEntity(int(mapWidth)/2, int(mapHeight)/2, "@", rl.White, "Player", true, RenderOrderActor, playerFighter, nil)

	entities := []*Entity{
		player,
	}

	gameMap := NewGameMap(mapWidth, mapHeight)
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

		var playerTurnResults []*ActionResult

		movement := action != nil && action.actionType == Movement
		if movement && gameState == PlayersTurn {
			destinationX := player.X + action.actionMovement.dx
			destinationY := player.Y + action.actionMovement.dy
			isBlocked := gameMap.IsPlayerBlocked(destinationX, destinationY)

			if !isBlocked {
				target := GetBlockingEntitiesAtLocation(entities, destinationX, destinationY)
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
			if resultType == Message {
				fmt.Println(playerTurnResult.ActionMessage)
			}
			if resultType == Dead {
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
						if resultType == Message {
							fmt.Println(enemyTurnResult.ActionMessage)
						}
						if resultType == Dead {
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

		renderAll(dejaVuFont, entities, player, gameMap, screenWidth, screenHeight, panelX, panelY, panelWidth, panelHeight)
		fovRecompute = false

		rl.EndDrawing()
	}

	rl.UnloadFont(dejaVuFont)

	rl.CloseWindow()
}
