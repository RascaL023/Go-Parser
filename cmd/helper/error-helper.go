package helper

import (
	"fmt"
	"os"
)

func Fatal(err error, code int) int {
	fmt.Fprintln(os.Stderr, "error:", err)
	return code
}

