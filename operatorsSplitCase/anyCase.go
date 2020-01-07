package operatorsSplitCase

import (
	"errors"
	"fmt"
	"strings"

	"goCalculator/processor"
)

func splitWithAnyDelimeter(r rune) bool {
	return r == '+' || r == '-' || r == '/' || r == '*'
}

func AnyDelimeterCase(expression, delimeter string) (float64, float64) {
	userInput := strings.FieldsFunc(expression, splitWithAnyDelimeter)
	n1 := userInput[0]
	n2 := userInput[1]

	convertedNum1 := processor.TrimAndConvertString(n1)
	convertedNum2 := processor.TrimAndConvertString(n2)

	if delimeter == "*" {
		fmt.Println(convertedNum1 * convertedNum2)
	} else if delimeter == "+" {
		fmt.Println(convertedNum1 + convertedNum2)
	} else if delimeter == "-" {
		fmt.Println(convertedNum1 - convertedNum2)
	} else if delimeter == "/" {
		fmt.Println(convertedNum1 / convertedNum2)
	} else {
		errors.New("invalid operator")
	}

	return convertedNum1, convertedNum2
}
