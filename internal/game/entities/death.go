package entities

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/internal/game/state"
)

func KillPlayer(player *Entity) (string, state.GameState) {
	player.Char = "%"
	player.Color = rl.Red

	return "You died!", state.PlayerDead
}

func KillMonster(monster *Entity) string {
	deathMessage := fmt.Sprintf("%s is dead!", monster.Name)
	monster.Char = "%"
	monster.Color = rl.Red
	monster.Blocks = false
	monster.Fighter = nil
	monster.Ai = nil
	monster.Name = "remains of " + monster.Name
	monster.RenderOrder = RenderOrderCorpse

	return deathMessage
}
