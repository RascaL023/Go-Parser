package state

type State struct {
	Version int 	 `json:"version"`
	Theme		string `json:"theme"`
	Waybar  string `json:"waybar"`
}
