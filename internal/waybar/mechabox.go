package waybar

type Mechabox struct {
    Rosewater string `json:"rosewater"`
    Flamingo  string `json:"flamingo"`
    Pink      string `json:"pink"`
    Mauve     string `json:"mauve"`
    Red       string `json:"red"`
    Maroon    string `json:"maroon"`
    Peach     string `json:"peach"`
    Yellow    string `json:"yellow"`
    Green     string `json:"green"`
    Teal      string `json:"teal"`
    Sky       string `json:"sky"`
    Sapphire  string `json:"sapphire"`
    Blue      string `json:"blue"`
    Lavender  string `json:"lavender"`

    Text      string `json:"text"`
    Subtext1  string `json:"subtext1"`
    Subtext0  string `json:"subtext0"`

    Overlay2  string `json:"overlay2"`
    Overlay1  string `json:"overlay1"`
    Overlay0  string `json:"overlay0"`

    Surface2  string `json:"surface2"`
    Surface1  string `json:"surface1"`
    Surface0  string `json:"surface0"`

    Base      string `json:"base"`
    Mantle   string `json:"mantle"`
    Crust    string `json:"crust"`

    MainBr    string `json:"mainBr"`
    MainBg    string `json:"mainBg"`
    MainFg    string `json:"mainFg"`

    Accent    string `json:"accent"`
    HoverBg   string `json:"hoverBg"`
    HoverFg   string `json:"hoverFg"`
    Outline   string `json:"outline"`

    Workspaces string `json:"workspaces"`
    Temperature string `json:"temperature"`
    Memory      string `json:"memory"`
    Cpu         string `json:"cpu"`
    Time        string `json:"time"`
    Date        string `json:"date"`
    Tray        string `json:"tray"`
    Volume      string `json:"volume"`
    Backlight   string `json:"backlight"`
    Battery     string `json:"battery"`

    Warning   string `json:"warning"`
    Critical  string `json:"critical"`
    Charging  string `json:"charging"`
}

