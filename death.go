package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/entity"
	"github.com/kotojo/roguelike_go/render_order"
)

func killPlayer(player *entity.Entity) (string, GameState) {
	player.Char = "%"
	player.Color = rl.Red

	return "You died!", PlayerDead
}

func killMonster(monster *entity.Entity) string {
	deathMessage := fmt.Sprintf("%s is dead!", monster.Name)
	monster.Char = "%"
	monster.Color = rl.Red
	monster.Blocks = false
	monster.Fighter = nil
	monster.Ai = nil
	monster.Name = "remains of " + monster.Name
	monster.RenderOrder = render_order.Corpse

	return deathMessage
}
