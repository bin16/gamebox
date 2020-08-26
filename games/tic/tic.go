package tic

type side int
type board [boardSize * boardSize]int
type history [][3]int

// Constants
const (
	boardSize   = 3
	lenOfWin    = 3
	right       = 1
	bottom      = boardSize
	bottomRight = boardSize + 1
	bottomLeft  = boardSize - 1

	Blank = iota
	SideX
	SideO

	StatusOpen
	StatusStarted
	StatusDraw
	StatusEnd

	OK
	NoChairs
	AlreadyIn
	NotStarted
	NotYourTurn
	NotFreeCell
)

func reverseOf(side int) (otherSide int) {
	if side == SideX {
		return SideO
	}

	return SideX
}

func nextSide(h history) int {
	hl := h[len(h)-1]
	sl := hl[0]
	return reverseOf(sl)
}

func testStep(b board, h history, s int, p, n int) (actionResult int) {
	if s != StatusStarted {
		return NotStarted
	}

	if ns := nextSide(h); ns != p {
		return NotYourTurn
	}

	if c := b[n]; c != Blank {
		return NotFreeCell
	}

	return OK
}

func commitStep(b board, h history, p, n int) (b1 board, h1 history, gameStatus, nextOrWinnerSide int) {
	b[n] = p
	h = append(h, [3]int{p, n, 0})

	status, winner := checkBoard(b, p)
	if status == StatusEnd {
		return b, h, status, winner
	}

	return b, h, status, reverseOf(p)
}

func checkBoard(b board, p int) (status, side int) {
	l := [][2]int{}
	for i := 0; i < boardSize; i++ {
		l = append(l, [2]int{i * bottom, right})                      // r-y
		l = append(l, [2]int{i * right, bottomRight})                 // br-x
		l = append(l, [2]int{i * bottom, bottomRight})                // br-y
		l = append(l, [2]int{i * right, bottom})                      // b-x
		l = append(l, [2]int{i * (boardSize - 1), bottomLeft})        // bl-x
		l = append(l, [2]int{(boardSize - 1) + i*bottom, bottomLeft}) // bl-y
	}

	for _, d := range l {
		line := pickLine(b, d[0], d[1])

		if xWin, _ := checkLine(line, SideX, lenOfWin); xWin {
			return StatusEnd, SideX
		}

		if oWin, _ := checkLine(line, SideO, lenOfWin); oWin {
			return StatusEnd, SideO
		}
	}

	bCount := 0
	for _, c := range b {
		if c == Blank {
			bCount++
		}
	}
	if bCount > 0 {
		return StatusStarted, Blank
	}

	return StatusEnd, Blank
}

func checkLine(lineOfSides []int, expectedSide int, more int) (found bool, index [2]int) {
	if len(lineOfSides) < more {
		return false, [2]int{0, 0}
	}

	var l, r int
	for i := 0; i < len(lineOfSides); i++ {
		if lineOfSides[i] != expectedSide {
			if r-l >= more {
				return true, [2]int{l, r}
			}

			l = i + 1
			r = l
		} else {
			r++
			if r == len(lineOfSides) && r-l >= more {
				return true, [2]int{l, r}
			}
		}
	}

	return false, [2]int{0, 0}
}

func pickLine(b board, n0, d int) []int {
	l := []int{}
	n1 := n0
	for i := 0; i < boardSize; i++ {
		l = append(l, b[n1])
		n1 += d
		if n1 >= len(b) {
			break
		}
		if d == bottomRight && n1/boardSize-(n1-d)/boardSize > 1 {
			break
		}
		if d == bottomLeft && n1%boardSize > (n1-d)%boardSize {
			break
		}
	}

	return l
}
