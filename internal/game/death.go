package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func killPlayer(player *Entity) (string, GameState) {
	player.Char = "%"
	player.Color = rl.Red

	return "You died!", PlayerDead
}

func killMonster(monster *Entity) string {
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
