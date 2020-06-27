package game

type BasicMonster struct {
	Owner *Entity
}

type Map interface {
	MapIsInFov(x, y int) bool
	Neighbors(x, y int) [][]int
	Cost(x, y, xx, yy int) int
}

func (b *BasicMonster) TakeTurn(target *Entity, gameMap Map) []*ActionResult {
	var results []*ActionResult
	monster := b.Owner
	if gameMap.MapIsInFov(monster.X, monster.Y) {
		if monster.DistanceTo(target) >= 2 {
			monster.MoveAStar(gameMap, target)
		} else if target.Fighter != nil && target.Fighter.Hp > 0 {
			attackResults := monster.Fighter.Attack(target)
			results = append(results, attackResults...)
		}
	}
	return results
}
