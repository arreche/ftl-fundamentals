// Package calculator provides math operations.
package calculator

import "errors"

// Add returns the sum of two integer numbers.
func Add(a, b int) int {
	return a + b
}

// Substract returns the substraction of two integer numbers.
func Substract(a, b int) int {
	return a - b
}

// Multiply returns the multiplication of two integer numbers.
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the division of two integer numbers.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b, nil
}
