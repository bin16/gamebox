package reversi

func nextSide(b board, h history) int {
	if len(h) == 0 {
		return SideBlack
	}

	ls := h[len(h)-1][0]
	rs := reverseOf(ls)
	if len(getLiveCells(b, rs)) > 0 {
		return rs
	} else if len(getLiveCells(b, ls)) > 0 {
		return ls
	}

	return Blank
}

func checkLine(l []int) bool {
	l1 := l[1:]
	s0 := l[0]
	num := 0
	for _, s := range l1 {
		if s == reverseOf(s0) {
			num++
		} else if s == s0 {
			return num > 0
		} else if s == Blank {
			return false
		}
	}

	return false
}

func pickLine(b board, n0, d int) []int {
	l := []int{}
	r0 := n0 / boardSize
	c0 := n0 % boardSize
	n1 := n0
	for i := 0; i < boardSize; i++ {
		l = append(l, b[n1])
		n1 += d
		if n1 >= len(b) || n1 < 0 {
			break
		}
		if (d == top || d == bottom) && n1%boardSize != c0 {
			break
		}
		if (d == right || d == left) && n1/boardSize != r0 {
			break
		}
		if (d == bottomRight || d == bottomLeft) && n1/boardSize != (n1-d)/boardSize+1 {
			break
		}
		if (d == topRight || d == topLeft) && n1/boardSize != (n1-d)/boardSize-1 {
			break
		}
	}

	return l
}

func cellIsLive(b board, index, side int) bool {
	if b[index] != Blank {
		return false
	}

	dl := []int{
		top, topRight, right,
		bottomRight, bottom,
		bottomLeft, left, topLeft,
	}
	for _, d := range dl {
		l := pickLine(b, index, d)
		l[0] = side
		if checkLine(l) {
			return true
		}
	}

	return false
}

func getLiveCells(b board, side int) []int {
	results := []int{}
	for i := range b {
		if cellIsLive(b, i, side) {
			results = append(results, i)
		}
	}

	return results
}

func reverseOf(side int) (otherSide int) {
	if side == SideBlack {
		return SideWhite
	}

	return SideBlack
}
