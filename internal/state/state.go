package state

type Theme struct {
	Name 						string `json:"name"`
	WallpaperPath		string `json:"wallpaperPath"`
}

type State struct {
	Version int 	 `json:"version"`
	Theme		Theme  `json:"theme"`
	Waybar  string `json:"waybar"`
}
