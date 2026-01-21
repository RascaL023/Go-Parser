package resolver

import (
	input "parsertry/internal/models/input/terminal"
	output "parsertry/internal/models/output/terminal"
)

func ResolveTerminal(in input.Terminal) output.TerminalOutput {
	return output.TerminalOutput {
		Foreground: in.Foreground,
		Background: in.Background,
		Regular: in.Palette.Normal,
		Bright: in.Palette.Bright,
	}
}

func ResolveKitty(out output.TerminalOutput, in input.Terminal) output.KittyOutput {
	return output.KittyOutput {
		TerminalOutput: out,
		SelectBackground: in.Selection.Background,
		SelectForeground: in.Selection.Foreground,
		Opacity: in.Opacity,
		Cursor: in.Cursor,
	}
}

func ResolveFoot(out output.TerminalOutput) output.FootOutput {
	return output.FootOutput {
		TerminalOutput: out,
	}
}
