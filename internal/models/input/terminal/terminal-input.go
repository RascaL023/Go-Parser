package input

type Palette struct {
	Normal 				[]string `json:"normal"`
	Bright 				[]string `json:"bright"`
}

type Terminal struct {
	Foreground		string `json:"foreground"`
	Background		string `json:"background"`

	Palette Palette `json:"palette"`

	Cursor				string 	`json:"cursor"`
	Opacity				float64 `json:"opacity"`

	Selection struct {
		Foreground  string `json:"foreground"`
		Background  string `json:"background"`
	} `json:"selection"`
}

