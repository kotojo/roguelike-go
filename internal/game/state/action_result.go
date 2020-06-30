package state

type ActionResultType int

const (
	Message ActionResultType = iota + 1
	Dead
)

type ActionResult struct {
	ResultType    ActionResultType
	ActionMessage string
}
