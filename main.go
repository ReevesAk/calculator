package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"goCalculator/operatorsSplitCase"
)

func main() {
	var input1, input2 int

	fmt.Println("***Welcome to Golang Calculator***")
	userInput := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your operation: ")
	userInput.Scan()
	userResponse1 := userInput.Text()

	if strings.ContainsAny(userResponse1, "+") {
		operatorsSplitCase.AnyDelimeterCase(userResponse1, "+")
	} else if strings.ContainsAny(userResponse1, "-") {
		operatorsSplitCase.AnyDelimeterCase(userResponse1, "-")
	} else if strings.ContainsAny(userResponse1, "/") {
		operatorsSplitCase.AnyDelimeterCase(userResponse1, "/")
	} else if strings.ContainsAny(userResponse1, "*") {
		operatorsSplitCase.AnyDelimeterCase(userResponse1, "*")
		fmt.Println(input1 * input2)
	}

}
