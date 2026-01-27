package cava

import (
	"encoding/json"
	"parsertry/internal/register"
	"parsertry/internal/renderer"
)

type CavaProcessor struct {}

func (CavaProcessor) Name() string { return "cava"; }

func (CavaProcessor) Resolve(input any) (any, error) {
	cava := input.(CavaInput);
	return ResolveCava(cava);
}

func (CavaProcessor) Parse(raw json.RawMessage) (any, error) {
	var cfg CavaInput;
	err := json.Unmarshal(raw, &cfg);
	return cfg, err;
}

func (CavaProcessor) Render(inputPath, outputPath string, data any) error {
	return renderer.Render(inputPath, outputPath, data);
}

func init() { register.Register(CavaProcessor{}); }

