package state

type ActionResultType int

const (
	ActionResultTypeMessage ActionResultType = iota + 1
	ActionResultTypeDead
)

type ActionResult struct {
	ResultType    ActionResultType
	ActionMessage Message
}
