package entities

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kotojo/roguelike_go/internal/game/state"
)

type Fighter struct {
	MaxHp   int
	Hp      int
	Defense int
	Power   int
	Owner   *Entity
}

func (f *Fighter) TakeDamage(amount int) []*state.ActionResult {
	var fightResult []*state.ActionResult
	f.Hp -= amount
	if f.Hp <= 0 {
		fightResult = append(fightResult, &state.ActionResult{
			ResultType: state.ActionResultTypeDead,
		})
	}
	return fightResult
}

func (f *Fighter) Attack(target *Entity) []*state.ActionResult {
	var fightResult []*state.ActionResult
	damage := f.Power - target.Fighter.Defense
	if damage > 0 {
		damageResult := target.Fighter.TakeDamage(damage)
		fightResult = append(fightResult, &state.ActionResult{
			ResultType:    state.ActionResultTypeMessage,
			ActionMessage: state.Message{Text: fmt.Sprintf("%s attacks %s for %d hit points.", f.Owner.Name, target.Name, damage), Color: rl.Black},
		})
		fightResult = append(fightResult, damageResult...)
	} else {
		fightResult = append(fightResult, &state.ActionResult{
			ResultType:    state.ActionResultTypeMessage,
			ActionMessage: state.Message{Text: fmt.Sprintf("%s attacks %s but does no damage.", f.Owner.Name, target.Name), Color: rl.Black},
		})
	}
	return fightResult
}
