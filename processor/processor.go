package processor

import (
	"strconv"
	"strings"
)

func TrimAndConvertString(takesString string) float64 {
	trimString := strings.Trim(takesString, " ")
	str, _ := strconv.ParseFloat(trimString, 32)
	return str
}
