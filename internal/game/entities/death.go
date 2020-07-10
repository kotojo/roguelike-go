package entities

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/internal/game/state"
)

func KillPlayer(player *Entity) (state.Message, state.GameState) {
	player.Char = "%"
	player.Color = rl.Red

	return state.Message{Text: "You died!", Color: rl.Red}, state.PlayerDead
}

func KillMonster(monster *Entity) state.Message {
	deathMessage := fmt.Sprintf("%s is dead!", monster.Name)
	monster.Char = "%"
	monster.Color = rl.Red
	monster.Blocks = false
	monster.Fighter = nil
	monster.Ai = nil
	monster.Name = "remains of " + monster.Name
	monster.RenderOrder = RenderOrderCorpse

	return state.Message{Text: deathMessage, Color: rl.Red}
}
