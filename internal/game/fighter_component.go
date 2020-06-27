package game

import "fmt"

type Fighter struct {
	MaxHp   int
	Hp      int
	Defense int
	Power   int
	Owner   *Entity
}

func (f *Fighter) TakeDamage(amount int) []*ActionResult {
	var fightResult []*ActionResult
	f.Hp -= amount
	if f.Hp <= 0 {
		fightResult = append(fightResult, &ActionResult{
			ResultType: Dead,
			DeadEntity: f.Owner,
		})
	}
	return fightResult
}

func (f *Fighter) Attack(target *Entity) []*ActionResult {
	var fightResult []*ActionResult
	damage := f.Power - target.Fighter.Defense
	if damage > 0 {
		damageResult := target.Fighter.TakeDamage(damage)
		fightResult = append(fightResult, &ActionResult{
			ResultType:    Message,
			ActionMessage: fmt.Sprintf("%s attacks %s for %d hit points.", f.Owner.Name, target.Name, damage),
		})
		fightResult = append(fightResult, damageResult...)
	} else {
		fightResult = append(fightResult, &ActionResult{
			ResultType:    Message,
			ActionMessage: fmt.Sprintf("%s attacks %s but does no damage.", f.Owner.Name, target.Name),
		})
	}
	return fightResult
}
