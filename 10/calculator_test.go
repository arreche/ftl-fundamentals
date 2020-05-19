package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 2},
		{a: -1, b: -2, want: -3},
		{a: 1, b: -2, want: -1},
		{a: 0, b: -1, want: -1},
		{a: 0, b: 0, want: 0},
	}
	for _, testCase := range testCases {
		got := calculator.Add(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 0},
		{a: -1, b: -2, want: 1},
		{a: -1, b: 2, want: -3},
		{a: 1, b: 0, want: 1},
		{a: 0, b: 0, want: 0},
	}
	for _, testCase := range testCases {
		got := calculator.Substract(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Subtract(%d, %d) want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 2, b: 3, want: 6},
		{a: 2, b: 1, want: 2},
		{a: 2, b: 0, want: 0},
		{a: 2, b: -1, want: -2},
		{a: -2, b: -2, want: 4},
	}
	for _, testCase := range testCases {
		got := calculator.Multiply(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Multiply(%d, %d) want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b        int
		want        int
		errExpected bool
	}{
		{a: 4, b: 2, want: 2},
		{a: 2, b: 2, want: 1},
		{a: 2, b: -1, want: -2},
		{a: -2, b: -1, want: 2},
		{a: -2, b: 1, want: -2},
		{a: 1, b: 0, want: 0, errExpected: true},
	}
	for _, testCase := range testCases {
		got, err := calculator.Divide(testCase.a, testCase.b)
		if testCase.errExpected && err == nil || !testCase.errExpected && err != nil {
			t.Errorf("Divide(%d, %d) expected error", testCase.a, testCase.b)
		}
		if testCase.want != got {
			t.Errorf("Divide(%d, %d) want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}
