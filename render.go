package main

import (
	"fmt"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/entity"
	"github.com/kotojo/roguelike_go/map_objects"
)

func renderAll(font rl.Font, entities []*entity.Entity, player *entity.Entity, gameMap *map_objects.GameMap, screenHeight int32) {
	for y := 0; y < gameMap.Height; y++ {
		for x := 0; x < gameMap.Width; x++ {
			tile := gameMap.Tiles[x][y]
			visible := gameMap.MapIsInFov(x, y)
			wall := false
			if tile != nil && tile.Blocked {
				wall = true
			}
			blockLoc := rl.Vector2{
				X: float32(x) * BlockSize,
				Y: float32(y) * BlockSize,
			}
			blockSize := rl.Vector2{
				X: BlockSize,
				Y: BlockSize,
			}
			color := rl.Black
			if visible {
				if wall {
					color = rl.Brown
				} else {
					color = rl.Gold
				}
				tile.Explored = true
			} else if tile.Explored {
				if wall {
					color = rl.DarkBlue
				} else {
					color = rl.Blue
				}
			}
			rl.DrawRectangleV(blockLoc, blockSize, color)
		}
	}

	sort.Slice(entities, func(i, j int) bool { return entities[i].RenderOrder < entities[j].RenderOrder })

	for _, entity := range entities {
		drawEntity(font, entity, gameMap)
	}
	rl.DrawTextEx(font, fmt.Sprintf("HP: %d/%d", player.Fighter.Hp, player.Fighter.MaxHp), rl.Vector2{X: 1, Y: float32(screenHeight) - BlockSize}, BlockSize, 0, rl.Black)
}

func drawEntity(font rl.Font, entity *entity.Entity, gameMap *map_objects.GameMap) {
	if gameMap.MapIsInFov(entity.X, entity.Y) {
		rec := rl.Rectangle{
			X:      float32(entity.X) * BlockSize,
			Y:      float32(entity.Y) * BlockSize,
			Width:  BlockSize,
			Height: BlockSize,
		}
		rl.DrawRectangleRec(rec, rl.Gold)
		rl.DrawTextRec(
			font,
			entity.Char,
			rec,
			BlockSize,
			0,
			false,
			entity.Color)
	}
}
