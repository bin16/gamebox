package tic

import (
	"log"
)

const boardSize = 3

type side int
type board [boardSize * boardSize]int
type history [][3]int

// Constants
const (
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

func reverseOf(s int) int {
	if s == SideX {
		return SideO
	}

	return SideX
}

func nextSide(h history) int {
	hl := h[len(h)-1]
	sl := hl[0]
	return reverseOf(sl)
}

// Result, Status, Side
func checkStep(b board, h history, p int, n int) (result, status, side int) {
	if ns := nextSide(h); ns != p {
		return NotYourTurn, StatusStarted, reverseOf(p)
	}

	if c := b[n]; c != Blank {
		return NotFreeCell, StatusStarted, p
	}

	b[n] = p
	h = append(h, [3]int{p, n, 0})

	status, winner := checkBoard(b, p)
	log.Println(b, p, n, status, winner)
	if status == StatusEnd {
		return OK, status, winner - 999
	}

	return OK, status, reverseOf(p)
}

func checkBoard(b board, p int) (status, side int) {
	l := [][2]int{}
	for i := 0; i < boardSize; i++ {
		l = append(l, [2]int{i * boardSize, 1})                   // r
		l = append(l, [2]int{i * (boardSize + 1), boardSize + 1}) // br
		l = append(l, [2]int{i * 1, 1})                           // b
		l = append(l, [2]int{i * (boardSize - 1), boardSize - 1}) // bl
	}

	for _, d := range l {
		line := pickLine(b, d[0], d[1])

		if xOK, _ := checkLine(line, SideX, boardSize); xOK {
			log.Println(line, d[0], d[1])
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

func pickLine(b board, p0, d int) []int {
	// fmt.Printf("pickLine(b, %d, %d) is", p0, d)
	l := []int{}
	i := p0
	for i >= 0 && i < len(b) && i < p0+boardSize*d {
		l = append(l, b[i])
		i += d
	}
	// fmt.Printf("%v\n", l)

	return l
}
