package theme

import (
	"encoding/json"
	_ "parsertry/internal/tools/cava"
	_ "parsertry/internal/tools/wlogout"
)

type Theme struct {
	Theme map[string]json.RawMessage `json:"config"`
}
