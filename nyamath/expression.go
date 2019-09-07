package nyamath

import (
	"errors"
	"strconv"
)

// Expression stores the infix, sufix and result of a math expression.
type Expression struct {
	infix  []interface{}
	suffix []interface{}
	result int
}

// Result returns the result of the expression e.
func (e Expression) Result() int {
	return e.result
}

func (e *Expression) suffix2Result() error {
	var s []interface{}

	for _, v := range e.suffix {
		switch v.(type) {
		case int:
			s = append(s, v)
		case string:
			if !IsUnary(v.(string)) {
				if len(s) < 2 {
					return errors.New("Invalid expression")
				}
				b := s[len(s)-1]
				s = s[:len(s)-1]
				a := s[len(s)-1]
				s = s[:len(s)-1]
				c, err := Calculate(a.(int), v.(string), b.(int))
				if err != nil {
					return errors.New(err.Error())
				}
				s = append(s, c)
			} else {
				if len(s) < 1 {
					return errors.New("Invalid expression")
				}
				top := s[len(s)-1]
				s = s[:len(s)-1]
				top, _ = CalculateU(v.(string), top.(int))
				s = append(s, top)
			}
		}
	}

	ans := s[len(s)-1]
	s = s[:len(s)-1]
	e.result = ans.(int)
	return nil
}

func (e *Expression) infix2Suffix() error {
	var s []interface{}

	for _, v := range e.infix {
		switch v.(type) {
		case int:
			e.suffix = append(e.suffix, v)
		case string:
			if v == "(" {
				s = append(s, v)
			} else if v == ")" {
				if len(s) < 1 {
					return errors.New("err")
				}
				top := s[len(s)-1]
				s = s[:len(s)-1]
				for top != "(" {
					e.suffix = append(e.suffix, top)
					if len(s) < 1 {
						return errors.New("err")
					}
					top = s[len(s)-1]
					s = s[:len(s)-1]
				}
			} else {
				if len(s) == 0 {
					s = append(s, v)
				} else {
					p1 := Priority(v.(string))
					top := s[len(s)-1]
					s = s[:len(s)-1]
					p2 := Priority(top.(string))
					for len(s) > 0 && p1 < p2 {
						e.suffix = append(e.suffix, top)
						s = s[:len(s)-1]
						if len(s) < 1 {
							return errors.New("err")
						}
						top := s[len(s)-1]
						p2 = Priority(top.(string))
					}
					s = append(s, v)
				}
			}
		}
	}

	for len(s) > 0 {
		top := s[len(s)-1]
		s = s[:len(s)-1]
		e.suffix = append(e.suffix, top)
	}
	return nil
}

func (e *Expression) string2Infix(s string) error {
	snum := ""
	last := '#'

	for _, v := range s {
		if v == ' ' {
			continue
		}
		if IsNumber(v) {
			snum = snum + string(v)
		} else if Valid(string(v)) {
			if len(snum) > 0 {
				num, _ := strconv.Atoi(snum)
				e.infix = append(e.infix, int(num))
				snum = ""
			}
			if v != '(' && last != ')' && !IsNumber(last) {
				o := "u" + string(v)
				if !Valid(o) {
					return errors.New("Invalid operator")
				}
				e.infix = append(e.infix, o)
			} else {
				e.infix = append(e.infix, string(v))
			}
		} else {
			return errors.New("Invalid operator")
		}
		last = v
	}

	if len(snum) > 0 {
		num, _ := strconv.Atoi(snum)
		e.infix = append(e.infix, int(num))
	}
	return nil
}

// New generates a expression from the string s.
func New(s string) (Expression, error) {
	var e Expression
	if err := e.string2Infix(s); err != nil {
		return e, err
	}
	if err := e.infix2Suffix(); err != nil {
		return e, err
	}
	if err := e.suffix2Result(); err != nil {
		return e, err
	}
	return e, nil
}
