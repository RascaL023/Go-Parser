package loader

import (
	"bufio"
	"os"
	"parsertry/internal/helper"
	"parsertry/internal/state"
	"strings"
)

type ToolMap struct {
	Name 					string
	TemplatePath	string
	OutputPath		string
}

func LoadToolMap(path string, state *state.State) ([]ToolMap, error) {
	path = helper.ExpandPath(path, state);

	fileMap, err := os.Open(path);
	if err != nil {
		return nil, err;
	}
	defer fileMap.Close();

	var res []ToolMap;
	scanner := bufio.NewScanner(fileMap);

	for scanner.Scan() {
		line := helper.ExpandPath(scanner.Text(), state);
		parts := strings.Split(line, "|");

		if len(parts) != 3 {
			continue;
		}

		res = append(res, ToolMap{
			Name: 				parts[0],
			TemplatePath: parts[1],
			OutputPath: 	parts[2],
		});
	}

	return res, scanner.Err();
}
