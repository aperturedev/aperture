package app

type ProjectionState int

const (
	ProjectionRunningState ProjectionState = 5
	ProjectionFailingState ProjectionState = 15
)

type ProjectionConfig struct {
	Name        string       `yaml:"name"`
	Type        string       `yaml:"type"`
	NetAssembly string       `yaml:"assembly"`
	Projections []Projection `yaml:"projections"`
}

type Projection struct {
	ID            string
	State         ProjectionState
	CurrentOffset uint64
	TrackOffset   string    `yaml:"trackOffset"` // TODO
	TypeFullName  string    `yaml:"typeFullName"`
	TypeName      string    `yaml:"typeName"`
	Name          string    `yaml:"name"`
	Description   string    `yaml:"description"`
	Handlers      []Handler `yaml:"handlers"`
}

type Handler struct {
	TypeFullName string `yaml:"typeFullName"`
	TypeName     string `yaml:"typeName"`
}
