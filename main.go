package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"goCalculator/operatorsSplitCase"
	//"goCalculator/processor"
)

func main() {
	var input []int
	var sum int

	for {

		fmt.Println("***Welcome to Golang Calculator***")
		userInput := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter your operation: ")
		userInput.Scan()
		userResponse1 := userInput.Text()

		if strings.ContainsAny(userResponse1, "+") {
			input = operatorsSplitCase.AStringArrayToInt(userResponse1)
			for val := range input {
				sum += input[val]
			}
			fmt.Println(sum)
		} else if strings.ContainsAny(userResponse1, "-") {
			input = operatorsSplitCase.AStringArrayToInt(userResponse1)
			sum = input[0]

			otherNumbers := input[1:]
			for index := range otherNumbers {

				sum -= otherNumbers[index]
			}
			fmt.Println(sum)
		} else if strings.ContainsAny(userResponse1, "/") {
			input = operatorsSplitCase.AStringArrayToInt(userResponse1)
			obelus := input[0]
			followUpNumbers := input[1:]
			for index := range followUpNumbers {
				obelus /= followUpNumbers[index]
			}
			fmt.Println(obelus)
		} else if strings.ContainsAny(userResponse1, "*") {
			input = operatorsSplitCase.AStringArrayToInt(userResponse1)
			times := input[0]
			restOfTheNumbers := input[1:]
			for index := range restOfTheNumbers {
				times *= restOfTheNumbers[index]
			}
			fmt.Println(times)
		}

		if userResponse1 == "" {
			break
		}

	}
}
