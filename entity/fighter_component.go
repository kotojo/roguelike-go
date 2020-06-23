package entity

type Fighter struct {
	MaxHp   int
	Hp      int
	Defense int
	Power   int
	Owner   *Entity
}
