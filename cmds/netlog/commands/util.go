package commands

import (
	"fmt"
	"os"
)

func fail(message string, err error) {
	_, _ = fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
}
