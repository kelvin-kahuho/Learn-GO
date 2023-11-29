package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter an expression (or 'exit' to quit): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Trim any leading or trailing whitespaces
		input = strings.TrimSpace(input)

		// Check for the exit command
		if input == "exit" {
			fmt.Println("Exiting calculator.")
			return
		}

		// Parse the expression and perform the calculation
		result, err := calculate(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
}

func calculate(expression string) (float64, error) {
	// Split the expression into operands and operator
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid expression format")
	}

	operand1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid operand: %v", err)
	}

	operator := parts[1]

	operand2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid operand: %v", err)
	}

	// Perform the calculation based on the operator
	var result float64
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		result = operand1 / operand2
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}

	return result, nil
}
