package tic

import (
	"testing"
)

func TestPickLine(t *testing.T) {
	// O _ O
	// _ X _
	// _ _ X
	var b board = [boardSize * boardSize]int{
		SideO, Blank, SideO, Blank, SideX, Blank, Blank, Blank, SideX,
	}
	tPickLine(t, b, 0, bottom)([]int{SideO, Blank, Blank})
	tPickLine(t, b, 1, bottom)([]int{Blank, SideX, Blank})
	tPickLine(t, b, 2, bottom)([]int{SideO, Blank, SideX})

	tPickLine(t, b, 0, bottomRight)([]int{SideO, SideX, SideX})
	tPickLine(t, b, 0+bottomRight, bottomRight)([]int{SideX, SideX})
	tPickLine(t, b, 0+bottomRight*2, bottomRight)([]int{SideX})
	tPickLine(t, b, 0+right, bottomRight)([]int{Blank, Blank})
	tPickLine(t, b, 0+right*2, bottomRight)([]int{SideO})
	tPickLine(t, b, 0+bottom, bottomRight)([]int{Blank, Blank})
	tPickLine(t, b, 0+bottom*2, bottomRight)([]int{Blank})

	tPickLine(t, b, 2, bottomLeft)([]int{SideO, SideX, Blank})
	tPickLine(t, b, 2+bottomLeft, bottomLeft)([]int{SideX, Blank})
	tPickLine(t, b, 2+bottomLeft*2, bottomLeft)([]int{Blank})
	tPickLine(t, b, 2-right, bottomLeft)([]int{Blank, Blank})
	tPickLine(t, b, 2-right*2, bottomLeft)([]int{SideO})
	tPickLine(t, b, 2+bottom, bottomLeft)([]int{Blank, Blank})
	tPickLine(t, b, 2+bottom*2, bottomLeft)([]int{SideX})

	tPickLine(t, b, 0, bottom)([]int{SideO, Blank, Blank})
	tPickLine(t, b, 0+right, bottom)([]int{Blank, SideX, Blank})
	tPickLine(t, b, 0+right*2, bottom)([]int{SideO, Blank, SideX})
}

func tPickLine(t *testing.T, b board, n0, d int) func(l0 []int) {
	return func(l0 []int) {
		t.Helper()
		if l1 := pickLine(b, n0, d); len(l1) != len(l0) {
			t.Errorf("Failed: pickLine(b, %d, %d), got %v, want %v", n0, d, l1, l0)
		} else {
			for i, c1 := range l1 {
				if c1 != l0[i] {
					t.Errorf("Failed: pickLine(b, %d, %d), got %v, want %v", n0, d, l1, l0)
					return
				}
			}
		}
	}
}

func TestMethods(t *testing.T) {
	// O _ O
	// _ X _
	// _ _ X
	var b board = [boardSize * boardSize]int{
		SideO, Blank, SideO, Blank, SideX, Blank, Blank, Blank, SideX,
	}
	var h history = [][3]int{
		[3]int{SideX, 4, 0}, // X
		[3]int{SideO, 2, 0}, // O
		[3]int{SideX, 8, 0}, // X
		[3]int{SideO, 0, 0}, // O
	}

	if nextSide(h) != SideX {
		t.Errorf("nextSide(h) should be %d, got %d", SideX, nextSide(h))
	}

	checkIt(t, b, h, SideX, 7)(OK, StatusStarted, SideO)
	checkIt(t, b, h, SideX, 6)(OK, StatusStarted, SideO)
	checkIt(t, b, h, SideX, 5)(OK, StatusStarted, SideO)
	checkIt(t, b, h, SideX, 4)(NotFreeCell, StatusStarted, SideX)
}

func checkIt(t *testing.T, b board, h history, p, n int) func(r0, s0, n0 int) {
	return func(r0, s0, n0 int) {
		t.Helper()
		r1, s1, n1 := checkStep(b, h, p, n)
		if r1 != r0 || s1 != s0 || n1 != n0 {
			t.Errorf("Failed, checkStep(%v, h, %d, %d) got (%d, %d, %d); want (%d, %d, %d)", b, p, n, r1, s1, n1, r0, s0, n0)
		}
	}
}

func TestCheckBoard(t *testing.T) {
	var b board = [boardSize * boardSize]int{
		SideO, Blank, SideO, Blank, SideX, Blank, Blank, Blank, SideX,
	}
	if s, ws := checkBoard(b, SideX); s != StatusStarted || ws != Blank {
		t.Errorf("Failed: checkBoard(b, %d), got (%d, %d), want (%d, %d)", SideX, s, ws, StatusStarted, Blank)
	}
	if s, ws := checkBoard(b, SideO); s != StatusStarted || ws != Blank {
		t.Errorf("Failed: checkBoard(b, %d), got (%d, %d), want (%d, %d)", SideO, s, ws, StatusStarted, Blank)
	}
}

func TestCheckLine(t *testing.T) {
	noResult := [2]int{}
	tCheckLine(t, []int{1, 2, 1, 2, 1, 2, 1}, 2, 3)(false, noResult)
	tCheckLine(t, []int{2, 0, 0}, 2, 3)(false, noResult)
	tCheckLine(t, []int{0, 2, 0}, 2, 3)(false, noResult)
	tCheckLine(t, []int{1, 0, 1}, 2, 3)(false, noResult)
	tCheckLine(t, []int{2, 2, 2}, 2, 3)(true, [2]int{0, 3})
	tCheckLine(t, []int{1, 2, 2, 2, 1}, 2, 3)(true, [2]int{1, 4})
	tCheckLine(t, []int{2, 0, 2, 0, 1, 0, 0, 0, 1}, 1, 3)(false, noResult)
	tCheckLine(t, []int{2, 0, 2, 0, 1, 0, 0, 0, 1}, 2, 3)(false, noResult)
	tCheckLine(t, []int{2, 0, 2, 0, 1, 0, 0, 0, 1}, 0, 3)(true, [2]int{5, 8})
	tCheckLine(t, []int{2, 0}, 2, 3)(false, noResult)
	tCheckLine(t, []int{2, 1}, 2, 3)(false, noResult)
	tCheckLine(t, []int{1, 1}, 1, 3)(false, noResult)
}

func tCheckLine(t *testing.T, l []int, want, more int) func(bool, [2]int) {
	return func(b0 bool, l0 [2]int) {
		t.Helper()
		if b1, l1 := checkLine(l, want, more); b1 != b0 {
			t.Errorf("Failed: checkLine(%v, %d, %d), got (%v, %v), want (%v, %v)", l, want, more, b1, l1, b0, l0)
		} else if len(l1) != len(l0) {
			t.Errorf("Failed: checkLine(%v, %d, %d), got (%v, %v), want (%v, %v)", l, want, more, b1, l1, b0, l0)
		} else {
			for i, a1 := range l1 {
				if a1 != l0[i] {
					t.Errorf("Failed: checkLine(%v, %d, %d), got (%v, %v), want (%v, %v)", l, want, more, b1, l1, b0, l0)
					return
				}
			}
		}
	}
}
