package processor

import (
	"strconv"
	"strings"
)

func TrimAndConvertString(takesString string) int {
	trimString := strings.Trim(takesString, " ")
	str, _ := strconv.Atoi(trimString)
	return str
}
