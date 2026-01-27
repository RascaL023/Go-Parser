package wlogout

type WlogoutInput struct {
	Window struct {
		Background   string		`json:"background"`
	} `json:"window"`

	Button struct {
		Background 	 string		`json:"background"`
		Focus struct {
			Background string		`json:"background"`
			Border 		 string 	`json:"border"`
		} `json:"focus"`
	} `json:"button"`
}
