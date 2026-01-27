package helper

type Palette struct {
	Normal []string `json:"normal"`
	Bright []string `json:"brignt"`
}

type Terminal struct {
	Foreground		string 		`json:"foreground"`
	Background		string 		`json:"background"`

	Cursor				string 		`json:"cursor"`
	Opacity				float64 	`json:"opacity"`

	Selection struct {
		Foreground		string 		`json:"foreground"`
		Background		string 		`json:"background"`
	} `json:"selection"`

	Palette Palette `json:"palette"`
}
