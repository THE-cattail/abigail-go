package nyamath

import (
	"errors"
	"fmt"

	"github.com/catsworld/random"
)

// Result here.
type Result struct {
	Min, Max, Value int
}

type operator struct {
	Op       rune
	LeftComb bool
}

// Expression stores the infix, sufix and result of a math expression.
type Expression struct {
	Infix  []interface{}
	Suffix []interface{}
	Result Result
}

var (
	priority = map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'%': 2,
		'^': 3,
		'd': 4,
	}
)

func isOperator(operator rune) bool {
	_, exist := priority[operator]
	return exist
}

func leftCombination(v rune, last interface{}) bool {
	if _, ok := last.(rune); ok {
		return false
	}
	return true
}

func calculate(x int, operator rune, y int) (int, error) {
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

func calculateMin(x int, operator rune, y int) (int, error) {
	if operator == 'd' {
		res1, err := calculate(x, '*', 1)
		if err != nil {
			return 0, err
		}
		res2, err := calculate(x, '*', y)
		if err != nil {
			return 0, err
		}
		return Min(res1, res2), nil
	}
	return calculate(x, operator, y)
}

func calculateMax(x int, operator rune, y int) (int, error) {
	if operator == 'd' {
		res1, err := calculate(x, '*', 1)
		if err != nil {
			return 0, err
		}
		res2, err := calculate(x, '*', y)
		if err != nil {
			return 0, err
		}
		return Max(res1, res2), nil
	}
	return calculate(x, operator, y)
}

func calculateU(unary rune, x int) (int, error) {
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

func (e *Expression) string2Infix(s string) error {
	nowNumber := 0
	nowNumberFlag := false
	for _, v := range s {
		if IsNumber(v) {
			if len(e.Infix) > 0 {
				if _, ok := e.Infix[len(e.Infix)-1].(int); ok {
					return errors.New("Invalid expression")
				}
			}
			nowNumber = nowNumber*10 + int(v-'0')
			nowNumberFlag = true
		} else {
			if nowNumberFlag {
				e.Infix = append(e.Infix, nowNumber)
				nowNumber = 0
				nowNumberFlag = false
			}
			if v != ' ' {
				e.Infix = append(e.Infix, v)
			}
		}
	}
	if nowNumberFlag {
		e.Infix = append(e.Infix, nowNumber)
	}
	return nil
}

func (e *Expression) infix2Suffix() error {
	stack := []*operator{}
	last := new(interface{})
	for _, symbol := range e.Infix {
		switch v := symbol.(type) {
		case int:
			e.Suffix = append(e.Suffix, v)
		case rune:
			if isOperator(v) {
				for len(stack) > 0 && ((leftCombination(v, last) && priority[v] <= priority[stack[len(stack)-1].Op]) || (!leftCombination(v, last) && priority[v] < priority[stack[len(stack)-1].Op])) {
					e.Suffix = append(e.Suffix, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, &operator{
					Op:       v,
					LeftComb: leftCombination(v, last),
				})
			} else if v == '(' {
				stack = append(stack, &operator{
					Op: v,
				})
			} else if v == ')' {
				for len(stack) > 0 && stack[len(stack)-1].Op != '(' {
					e.Suffix = append(e.Suffix, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				if len(stack) == 0 || stack[len(stack)-1].Op != '(' {
					return errors.New("Invalid expression")
				}
				stack = stack[:len(stack)-1]
			} else {
				return errors.New("Invalid expression")
			}
		}
	}
	for len(stack) > 0 {
		if stack[len(stack)-1].Op == '(' || stack[len(stack)-1].Op == ')' {
			return errors.New("Invalid expression")
		}
		e.Suffix = append(e.Suffix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return nil
}

func (e *Expression) suffix2Result() error {
	stack := []*Result{}
	for _, symbol := range e.Suffix {
		switch v := symbol.(type) {
		case int:
			stack = append(stack, &Result{
				Min:   v,
				Max:   v,
				Value: v,
			})
		case *operator:
			if (v.LeftComb && len(stack) < 2) || (!v.LeftComb && len(stack) < 1) {
				return errors.New("Invalid expression")
			}
			if v.LeftComb {
				res, err := calculate(stack[len(stack)-2].Value, v.Op, stack[len(stack)-1].Value)
				if err != nil {
					return fmt.Errorf("Calculating: %v", err)
				}
				res1, err := calculateMin(stack[len(stack)-2].Min, v.Op, stack[len(stack)-1].Min)
				if err != nil {
					return fmt.Errorf("Calculating: %v", err)
				}
				res2, err := calculateMin(stack[len(stack)-2].Min, v.Op, stack[len(stack)-1].Max)
				if err != nil {
					return fmt.Errorf("Calculating: %v", err)
				}
				res3, err := calculateMin(stack[len(stack)-2].Max, v.Op, stack[len(stack)-1].Min)
				if err != nil {
					return fmt.Errorf("Calculating: %v", err)
				}
				res4, err := calculateMin(stack[len(stack)-2].Max, v.Op, stack[len(stack)-1].Max)
				if err != nil {
					return fmt.Errorf("Calculating: %v", err)
				}
				resMin := Min(res1, res2, res3, res4)
				res1, err = calculateMax(stack[len(stack)-2].Min, v.Op, stack[len(stack)-1].Min)
				if err != nil {
					return fmt.Errorf("Calculating: %v", err)
				}
				res2, err = calculateMax(stack[len(stack)-2].Min, v.Op, stack[len(stack)-1].Max)
				if err != nil {
					return fmt.Errorf("Calculating: %v", err)
				}
				res3, err = calculateMax(stack[len(stack)-2].Max, v.Op, stack[len(stack)-1].Min)
				if err != nil {
					return fmt.Errorf("Calculating: %v", err)
				}
				res4, err = calculateMax(stack[len(stack)-2].Max, v.Op, stack[len(stack)-1].Max)
				if err != nil {
					return fmt.Errorf("Calculating: %v", err)
				}
				resMax := Max(res1, res2, res3, res4)
				stack = stack[:len(stack)-2]
				stack = append(stack, &Result{
					Min:   resMin,
					Max:   resMax,
					Value: res,
				})
			} else {
				res, err := calculateU(v.Op, stack[len(stack)-1].Value)
				if err != nil {
					return fmt.Errorf("Calculating: %v", err)
				}
				resMin := 0
				resMax := 0
				if v.Op == 'd' {
					resMin, err = calculateU('+', 1)
					resMax, err = calculateU('+', stack[len(stack)-1].Value)
				} else {
					resMin = res
					resMax = res
				}
				stack = stack[:len(stack)-1]
				stack = append(stack, &Result{
					Min:   resMin,
					Max:   resMax,
					Value: res,
				})
			}
		}
	}
	if len(stack) != 1 {
		return errors.New("Invalid expression")
	}
	e.Result = *stack[0]
	return nil
}

// New generates a expression from the string s.
func New(s string) (*Expression, error) {
	e := &Expression{}
	err := e.string2Infix(s)
	if err != nil {
		return nil, fmt.Errorf("Generating new expression: %v", err)
	}
	err = e.infix2Suffix()
	if err != nil {
		return nil, fmt.Errorf("Generating new expression: %v", err)
	}
	err = e.suffix2Result()
	if err != nil {
		return nil, fmt.Errorf("Generating new expression: %v", err)
	}
	return e, nil
}
