package service

import (
	"errors"
)

func Calculate(a, b int, op string) (int, error) {
	switch op {
	case "add":
		return a + b, nil
	case "sub":
		return a - b, nil
	case "mul":
		return a * b, nil

	case "div":
		if b == 0 {
			return 0, errors.New("cannot divide zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("operations Not supported")
	}
}
