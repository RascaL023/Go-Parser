package kitty

func BuildKitty(cfg KittyConfig) Kitty {
	p := cfg.Palette

	return Kitty{
		Cursor:   cfg.Cursor,
		Fg:       cfg.Fg,
		Bg:       cfg.Bg,
		SelectBg: cfg.Selection.Background,
		SelectFg: cfg.Selection.Foreground,
		Opacity:  cfg.Opacity,

		Color0:  p.Black,
		Color1:  p.Red,
		Color2:  p.Green,
		Color3:  p.Yellow,
		Color4:  p.Blue,
		Color5:  p.Magenta,
		Color6:  p.Cyan,
		Color7:  p.White,
		Color8:  p.BrightBlack,
		Color9:  p.BrightRed,
		Color10: p.BrightGreen,
		Color11: p.BrightYellow,
		Color12: p.BrightBlue,
		Color13: p.BrightMagenta,
		Color14: p.BrightCyan,
		Color15: p.BrightWhite,
	}
}
