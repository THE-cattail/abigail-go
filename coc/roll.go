package coc

import "github.com/catsworld/random"

// CheckResult here.
type CheckResult struct {
	Succ, Big, Level, N, Number int
}

var (
	// CheckSuccess here.
	CheckSuccess = 1
	// CheckFailure here.
	CheckFailure = -1
	// CheckBigSuccess here.
	CheckBigSuccess = 1
	// CheckBigFailure here.
	CheckBigFailure = -1
	// CheckHardSuccess here.
	CheckHardSuccess = 1
	// CheckVeryHardSuccess here.
	CheckVeryHardSuccess = 2
)

// Success here.
func (r CheckResult) Success() bool {
	return r.N > 0 && r.N < 100 && r.Succ == CheckSuccess
}

// Failure here.
func (r CheckResult) Failure() bool {
	return r.N > 0 && r.N < 100 && r.Succ == CheckFailure
}

// BigSuccess here.
func (r CheckResult) BigSuccess() bool {
	return r.Big == CheckBigSuccess
}

// BigFailure here.
func (r CheckResult) BigFailure() bool {
	return r.Big == CheckBigFailure
}

// HardSuccess here.
func (r CheckResult) HardSuccess() bool {
	return r.Level == CheckHardSuccess
}

// VeryHardSuccess here.
func (r CheckResult) VeryHardSuccess() bool {
	return r.Level == CheckVeryHardSuccess
}

// Check here.
func Check(n int) CheckResult {
	result := CheckResult{
		N:      n,
		Number: random.Int(1, 100),
	}
	if result.Number <= n {
		result.Succ = CheckSuccess
	} else {
		result.Succ = CheckFailure
	}
	if result.Number <= n/2 {
		result.Level = CheckHardSuccess
	}
	if result.Number <= n/5 {
		result.Level = CheckVeryHardSuccess
	}
	if n >= 50 {
		if result.Number == 100 {
			result.Big = CheckBigFailure
		}
	} else {
		if result.Number > 95 {
			result.Big = CheckBigFailure
		}
	}
	if result.Number == 1 {
		result.Big = CheckBigSuccess
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
