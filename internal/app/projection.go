package app

type ProjectionState int

const (
	ProjectionRunningState ProjectionState = 5
	ProjectionFailingState ProjectionState = 15
)

type Projection struct {
	Name        string
	Description string
	State       ProjectionState
}
