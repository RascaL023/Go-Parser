package wlogout

import (
	"encoding/json"
	"parsertry/internal/context"
	"parsertry/internal/register"
	"parsertry/internal/renderer"
)

type WlogoutProcessor struct {}

func (WlogoutProcessor) Name() string { return "wlogout"; }

func (WlogoutProcessor) Resolve(input any, _ *context.Context) (any, error) {
	wlogout := input.(WlogoutInput);
	return resolve(wlogout);
}

func (WlogoutProcessor) Parse(raw json.RawMessage) (any, error) {
	var cfg WlogoutInput;
	err := json.Unmarshal(raw, &cfg);
	return cfg, err;
}

func (WlogoutProcessor) Render(inputPath, outputPath string, data any) error {
	return renderer.Render(inputPath, outputPath, data);
}

func init() { register.RegisterTool(WlogoutProcessor{}); }
