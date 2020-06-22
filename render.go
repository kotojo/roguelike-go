package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/entity"
	"github.com/kotojo/roguelike_go/map_objects"
)

func renderAll(font rl.Font, entities []*entity.Entity, gameMap *map_objects.GameMap) {
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
	for _, entity := range entities {
		drawEntity(font, entity, gameMap)
	}
}

func drawEntity(font rl.Font, entity *entity.Entity, gameMap *map_objects.GameMap) {
	if gameMap.MapIsInFov(entity.X, entity.Y) {
		rl.DrawTextRec(
			font,
			entity.Char,
			rl.Rectangle{
				X:      float32(entity.X) * BlockSize,
				Y:      float32(entity.Y) * BlockSize,
				Width:  BlockSize,
				Height: BlockSize,
			},
			BlockSize,
			0,
			false,
			entity.Color)
	}
}
