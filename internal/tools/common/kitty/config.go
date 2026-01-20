package kitty

type KittyConfig struct {
	Cursor   string `json:"cursor"`
	Fg       string `json:"fg"`
	Bg       string `json:"bg"`

	Selection struct {
		Background string `json:"background"`
		Foreground string `json:"foreground"`
	} `json:"selection"`

	Opacity string `json:"opacity"`

	Palette struct {
		Black   string `json:"black"`
		Red     string `json:"red"`
		Green   string `json:"green"`
		Yellow  string `json:"yellow"`
		Blue    string `json:"blue"`
		Magenta string `json:"magenta"`
		Cyan    string `json:"cyan"`
		White   string `json:"white"`

		BrightBlack   string `json:"brightBlack"`
		BrightRed     string `json:"brightRed"`
		BrightGreen   string `json:"brightGreen"`
		BrightYellow  string `json:"brightYellow"`
		BrightBlue    string `json:"brightBlue"`
		BrightMagenta string `json:"brightMagenta"`
		BrightCyan    string `json:"brightCyan"`
		BrightWhite   string `json:"brightWhite"`
	} `json:"palette"`
}

