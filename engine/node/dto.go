package node

type Type string

var (
	TypeText   Type = "text"
	TypeModel  Type = "model"
	TypePlugin Type = "plugin"
)

type Node struct {
	ID                         string
	Stage                      string // trigger, execution, delivery
	Type                       Type
	Dependencies               []Dependency
	NextNode                   []NextNode
	Config                     Config
	TerminateWorkflowOnFailure bool
}

type Config struct {
	Text   *TextNodeConfig
	Plugin *PluginNodeConfig
	Model  *ModelNodeConfig
}

type ModelNodeConfig struct {
	ModelEndpoint string
	Temperature   float32
}

type TextNodeConfig struct {
	EndpointURL string
}

type PluginNodeConfig struct {
	EndpointURL    string
	Authentication Auth
}

type Auth struct {
	Username string
	Password string
}

type Dependency struct {
	NodeID string
}

type NextNode struct {
	NodeID string
}

type Output struct {
}
