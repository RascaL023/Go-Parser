package cava

func ResolveCava(input CavaInput) (Cava, error) {
	return Cava{
		Gradient1: input.GradientColor[0],
		Gradient2: input.GradientColor[1],
		Gradient3: input.GradientColor[2],
		Gradient4: input.GradientColor[3],
		Gradient5: input.GradientColor[4],
	}, nil;
}
