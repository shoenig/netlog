package netlog

import (
	"fmt"
	"strings"
)

func format(msg string, args ...interface{}) string {
	if len(args)%2 != 0 {
		return msg + " !BAD-ARGS! " + fmt.Sprintf("%v", args)
	}
	var sb strings.Builder
	sb.WriteString(msg)
	sb.WriteString(" ")
	for i := 0; i < len(args); i += 2 {
		sb.WriteString(fmt.Sprintf("%s=%v ", args[i].(string), args[i+1]))
	}
	return strings.TrimSpace(sb.String())
}
