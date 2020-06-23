package entity

import (
	"fmt"
)

type BasicMonster struct {
	Owner *Entity
}

func (b *BasicMonster) TakeTurn() {
	fmt.Printf("The %s wonders when it will get to move.\n", b.Owner.Name)
}
