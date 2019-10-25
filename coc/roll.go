package coc

import "github.com/catsworld/random"

// CheckResult here.
type CheckResult struct {
	Succ, Great, Level, N, Number int
}

var (
	// Succ means the check is successful.
	Succ = 1
	// Fail means the check is failed.
	Fail = -1
	// GreatSucc means the check is great successful.
	GreatSucc = 1
	// GreatFail means the check is great failed.
	GreatFail = -1
	// DiffSucc means the check is difficultly successful.
	DiffSucc = 1
	// ExDiffSucc means the check is extremely difficultly successful.
	ExDiffSucc = 2
)

// Check returns the result of a check.
func Check(n int) CheckResult {
	result := CheckResult{
		N:      n,
		Number: random.Int(1, 100),
	}

	if n <= 0 {
		result.Succ = Fail
		return result
	}

	if result.Number <= n {
		result.Succ = Succ
	} else {
		result.Succ = Fail
	}

	if result.Number <= n/2 {
		result.Level = DiffSucc
	}
	if result.Number <= n/5 {
		result.Level = ExDiffSucc
	}

	if n >= 50 && n < 100 {
		if result.Number == 100 {
			result.Great = GreatFail
		}
	} else if n < 50 {
		if result.Number > 95 {
			result.Great = GreatFail
		}
	}

	if result.Number == 1 {
		result.Great = GreatSucc
	}
	return result
}

// Roll returns the result of xdy defined in CoC rules.
func Roll(x, y int) int {
	ret := int(0)
	for i := int(0); i < x; i++ {
		ret += random.Int(1, y)
	}
	return ret
}
