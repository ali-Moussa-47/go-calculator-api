package tests

import (
	"calculatorApi/service"
	"testing"
)

func TestCalculateAll(t *testing.T) {
	tests := []struct {
		a, b       int
		op         string
		expected   int
		shouldFail bool
	}{
		{3, 2, "add", 5, false},
		{3, 2, "sub", 1, false},
		{4, 2, "mul", 8, false},
		{8, 2, "div", 4, false},
		{5, 0, "div", 0, true},
		{2, 2, "pow", 0, true},
	}

	for _, test := range tests {
		result, err := service.Calculate(test.a, test.b, test.op)
		if test.shouldFail && err == nil {
			t.Errorf("Expected error for op=%s, got none", test.op)
		}
		if !test.shouldFail && (err != nil || result != test.expected) {
			t.Errorf("Expected %d for op=%s, got %d (err=%v)",
				test.expected, test.op, result, err)
		}
	}
}
