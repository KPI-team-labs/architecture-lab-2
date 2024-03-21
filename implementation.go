package lab2

import (
	"fmt"
	"strings"
	"unicode"
)

// PrefixToInfix function is to convert prefix expression into infix
// Input accepts a string representation of prefix expression
// Error if the prefix expression is not valid or contains numerical characters
// Output if the prefix expression is appropriate string will be converted to infix
func PrefixToInfix(prefix string) (string, error) {
	availableOptions := map[string]bool{"+": true, "-": true, "*": true, "/": true, "^": true}
	stack := []string{}
	elements := strings.Split(prefix, " ")
	for i := len(elements) - 1; i >= 0; i-- {
		element := elements[i]
		if _, valid := availableOptions[element]; valid {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid prefix expression")
			}
			secondOperand := stack[len(stack)-2]
			firstOperand := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			expression := "(" + firstOperand + " " + element + " " + secondOperand + ")"
			stack = append(stack, expression)
		} else {
			stack = append(stack, element)
		}
		if unicode.IsLetter(rune(element[0])) {
			return "", fmt.Errorf("letters are not allowed")
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("your prefix expression is not valid")
	}
	return stack[0], nil
}
