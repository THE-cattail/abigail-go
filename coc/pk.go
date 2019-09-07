package coc

var (
	// PKAWin here.
	PKAWin = -1
	// PKDraw here.
	PKDraw = 0
	// PKBWin here.
	PKBWin = 1
)

// PK here.
func PK(a, b CheckResult) int {
	if a.BigSuccess() && b.BigSuccess() {
		return PKDraw
	}
	if a.BigSuccess() {
		return PKAWin
	}
	if b.BigSuccess() {
		return PKBWin
	}
	if a.Success() && b.Success() {
		if a.Level == b.Level {
			return PKDraw
		} else if a.Level > b.Level {
			return PKAWin
		} else {
			return PKBWin
		}
	}
	if a.Success() {
		return PKAWin
	}
	if b.Success() {
		return PKBWin
	}
	return PKDraw
}
