package wlogout

func resolve(in WlogoutInput) (Wlogout, error) {
	return Wlogout{
		WindowBg: in.Window.Background,

		ButtonBg: in.Button.Background,
		ButtonFocusBg: in.Button.Focus.Background,
		ButtonFocusBorder: in.Button.Focus.Border,
	}, nil;
}
