package render_order

type RenderOrder int

const (
	Corpse RenderOrder = iota
	Item
	Actor
)
