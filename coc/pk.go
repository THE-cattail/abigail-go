package coc

var (
	// PKAWin means that A is the winner in this PK.
	PKAWin = -1
	// PKDraw means that the result of this PK is a draw.
	PKDraw = 0
	// PKBWin means that B is the winner in this PK.
	PKBWin = 1
)

// PK returns the result of a PK.
func PK(a, b CheckResult) int {
	if a.Great == GreatSucc && b.Great == GreatSucc {
		return PKDraw
	}

	if a.Great == GreatSucc {
		return PKAWin
	}
	if b.Great == GreatSucc {
		return PKBWin
	}

	if a.Succ == Succ && b.Succ == Succ {
		if a.Level == b.Level {
			return PKDraw
		} else if a.Level > b.Level {
			return PKAWin
		} else {
			return PKBWin
		}
	}

	if a.Succ == Succ {
		return PKAWin
	}
	if b.Succ == Succ {
		return PKBWin
	}

	return PKDraw
}
