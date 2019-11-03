package nyamath

import (
	"errors"

	"github.com/catsworld/random"
)

var (
	priority = map[rune]int{
		'+': 2,
		'-': 2,
		'*': 3,
		'/': 3,
		'%': 3,
		'^': 4,
		'd': 5,
	}
)

// Operator checks if the operator is a valid operator.
func Operator(operator rune) bool {
	_, exist := priority[operator]
	return exist
}

// Priority returns the priority of the operator.
func Priority(operator rune) int {
	if ret, exist := priority[operator]; exist {
		return ret
	}
	return 0
}

// Calculate returns the result of calculating x and y with the operator.
func Calculate(x int, operator rune, y int) (int, error) {
	if operator == '+' {
		return x + y, nil
	} else if operator == '-' {
		return x - y, nil
	} else if operator == '*' {
		return x * y, nil
	} else if operator == '/' {
		if y == 0 {
			return 0, errors.New("Divide by 0")
		}
		return x / y, nil
	} else if operator == '%' {
		if y == 0 {
			return 0, errors.New("Divide by 0")
		}
		return x % y, nil
	} else if operator == 'd' {
		if x > 100 {
			return 0, errors.New("Too many times")
		}
		ret := 0
		for i := 0; i < x; i++ {
			ret += random.Int(1, y)
		}
		return ret, nil
	} else if operator == '^' {
		return Qpow(x, y), nil
	} else {
		return 0, errors.New("Invalid operator")
	}
}

// CalculateU returns the result of calculating a with the unary.
func CalculateU(unary rune, x int) (int, error) {
	if unary == '-' {
		return -x, nil
	} else if unary == '+' {
		return x, nil
	} else if unary == 'd' {
		return random.Int(1, x), nil
	} else {
		return 0, errors.New("Invalid unary")
	}
}
