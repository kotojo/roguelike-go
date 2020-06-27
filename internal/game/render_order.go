package game

type RenderOrder int

const (
	RenderOrderCorpse RenderOrder = iota
	RenderOrderItem
	RenderOrderActor
)
