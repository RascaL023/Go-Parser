package register

import "encoding/json"

var Registered = map[string]Processor{}

type Processor interface {
    Name() string
		Parse(json.RawMessage) (any, error)
    Resolve(input any) (any, error)
    Render(templatePath, outputPath string, data any) error
}

func Register(p Processor) {
    Registered[p.Name()] = p;
}

func Get(name string) (Processor, bool) {
    p, ok := Registered[name];
    return p, ok;
}
