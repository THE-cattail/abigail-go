package nyamath

import (
	"errors"

	"github.com/catsworld/random"
)

var (
	priority = map[string]int{
		"(":  1,
		"+":  2,
		"-":  2,
		"*":  3,
		"/":  3,
		"%":  3,
		"^":  4,
		"d":  5,
		"ud": 6,
		"u-": 7,
		"u+": 7,
		")":  8,
	}
)

// Valid checks if the operator is a valid operator.
func Valid(operator string) bool {
	_, exist := priority[operator]
	return exist
}

// Priority returns the priority of the operator.
func Priority(operator string) int {
	if ret, exist := priority[operator]; exist {
		return ret
	}
	return 0
}

// IsUnary checks if the operator is an unary.
func IsUnary(operator string) bool {
	if Valid(operator) && operator[0] == 'u' {
		return true
	}
	return false
}

// Calculate returns the result of calculating x and y with the operator.
func Calculate(x int, operator string, y int) (int, error) {
	if operator == "+" {
		return x + y, nil
	} else if operator == "-" {
		return x - y, nil
	} else if operator == "*" {
		return x * y, nil
	} else if operator == "/" {
		if y == 0 {
			return 0, errors.New("Divide by 0")
		}
		return x / y, nil
	} else if operator == "%" {
		if y == 0 {
			return 0, errors.New("Divide by 0")
		}
		return x % y, nil
	} else if operator == "d" {
		if x > 100 {
			return 0, errors.New("Too many times")
		}
		ret := 0
		for i := 0; i < x; i++ {
			ret += random.Int(1, y)
		}
		return ret, nil
	} else if operator == "^" {
		return Qpow(x, y), nil
	} else {
		return 0, errors.New("Invalid operator")
	}
}

// CalculateU returns the result of calculating a with the unary.
func CalculateU(unary string, x int) (int, error) {
	if unary == "u-" {
		return -x, nil
	} else if unary == "u+" {
		return x, nil
	} else if unary == "ud" {
		return random.Int(1, x), nil
	} else {
		return 0, errors.New("Invalid unary")
	}
}
