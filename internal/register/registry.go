package register

import (
	"encoding/json"
	"parsertry/internal/context"
)

var RegisteredTool = map[string]Tool{}

type Parser interface {
	Parse(json.RawMessage) (any, error)
}

type Resolver interface {
	Resolve(any, *context.Context) (any, error)
}

type Renderer interface {
	Render(templatePath, outputPath string, data any) error
}

type Tool interface {
    Name() string
		Parser
		Resolver
		// Renderer
}

func RegisterTool(t Tool) {
    RegisteredTool[t.Name()] = t;
}

func Get(name string) (Tool, bool) {
    p, ok := RegisteredTool[name];
    return p, ok;
}


