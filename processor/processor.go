package processor

import (
	"strings"
)

func TrimString(takesString string) string {
	return strings.ReplaceAll(takesString, " ", "")
}
