package theme

import (
	"encoding/json"
	_ "parsertry/internal/tools/cava"
)

type Theme struct {
	Theme map[string]json.RawMessage `json:"config"`
}
