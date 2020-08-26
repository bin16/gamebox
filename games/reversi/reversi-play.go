package reversi

func testStep(b board, h history, s int, p, n int) (actionResult int) {
	if s != StatusStarted {
		return NotStarted
	}

	if ns := nextSide(b, h); ns != p {
		return NotYourTurn
	}

	if c := b[n]; c != Blank {
		return NotFreeCell
	}

	if !cellIsLive(b, n, p) {
		return NotFreeCell
	}

	return OK
}

func commitStep(b board, h history, p, n int) (b1 board, h1 history, gameStatus, nextOrWinnerSide int) {
	b[n] = p
	h = append(h, [3]int{p, n, 0})

	ns := nextSide(b, h)
	if ns == Blank {
		ws := winnerOf(b)
		if ws == Blank {
			return b, h, StatusDraw, ws
		}
		return b, h, StatusEnd, ws
	}

	return b, h, StatusStarted, ns
}
