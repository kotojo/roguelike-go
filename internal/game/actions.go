package game

type ActionType int

const (
	Movement ActionType = iota + 1
	Fullscreen
)

type Action struct {
	actionType     ActionType
	actionMovement *ActionMovement
}

type ActionMovement struct {
	dx int
	dy int
}
