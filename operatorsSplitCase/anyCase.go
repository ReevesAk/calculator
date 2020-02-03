package operatorsSplitCase

import (
	//"errors"
	//"fmt"
	"strconv"
	"strings"

	"goCalculator/processor"
)

func splitWithAnyDelimeter(r rune) bool {
	return r == '+' || r == '-' || r == '/' || r == '*' || r == '(' || r == ')'
}

func AnyDelimeterCase(expression string) []string {
	userInput := strings.FieldsFunc(expression, splitWithAnyDelimeter)

	return userInput
}

func AStringArrayToInt(input string) []int {
	trimmedString := processor.TrimString(input)
	inputStr := AnyDelimeterCase(trimmedString)
	arr := make([]int, len(inputStr))
	for index, value := range inputStr {
		arr[index], _ = strconv.Atoi(value)
	}
	return arr
}
