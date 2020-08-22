package tic

type side int
type board [boardSize * boardSize]int
type history [][3]int

// Constants
const (
	boardSize   = 3
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

func checkStep(b board, h history, p int, n int) (actionResult, gameStatus, nextOrWinnerSide int) {
	if ns := nextSide(h); ns != p {
		return NotYourTurn, StatusStarted, reverseOf(p)
	}

	if c := b[n]; c != Blank {
		return NotFreeCell, StatusStarted, p
	}

	b[n] = p
	h = append(h, [3]int{p, n, 0})

	status, winner := checkBoard(b, p)
	if status == StatusEnd {
		return OK, status, winner
	}

	return OK, status, reverseOf(p)
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

		if xOK, _ := checkLine(line, SideX, boardSize); xOK {
			return StatusEnd, SideX
		}

		if oOK, _ := checkLine(line, SideO, boardSize); oOK {
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

func checkLine(line []int, want int, more int) (bool, []int) {
	if len(line) < more {
		return false, []int{}
	}

	var l, r int
	for i := 0; i < len(line); i++ {
		if line[i] != want {
			if r-l >= more {
				return true, line[l:r]
			}

			l = i + 1
			r = l
		} else {
			r++
			if r == len(line) && r-l >= more {
				return true, line[l:r]
			}
		}
	}

	return false, []int{}
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
