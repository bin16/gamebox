package tic

import (
	"testing"
)

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
	tCheckLine(t, []int{1, 2, 1, 2, 1, 2, 1}, 2, 3)(false, []int{})
	tCheckLine(t, []int{2, 0, 0}, 2, 3)(false, []int{})
	tCheckLine(t, []int{0, 2, 0}, 2, 3)(false, []int{})
	tCheckLine(t, []int{1, 0, 1}, 2, 3)(false, []int{})
	tCheckLine(t, []int{2, 2, 2}, 2, 3)(true, []int{2, 2, 2})
	tCheckLine(t, []int{1, 2, 2, 2, 1}, 2, 3)(true, []int{2, 2, 2})
	tCheckLine(t, []int{2, 0, 2, 0, 1, 0, 0, 0, 1}, 1, 3)(false, []int{})
	tCheckLine(t, []int{2, 0, 2, 0, 1, 0, 0, 0, 1}, 2, 3)(false, []int{})
	tCheckLine(t, []int{2, 0, 2, 0, 1, 0, 0, 0, 1}, 0, 3)(true, []int{0, 0, 0})
	tCheckLine(t, []int{2, 0}, 2, 3)(false, []int{})
	tCheckLine(t, []int{2, 1}, 2, 3)(false, []int{})
	tCheckLine(t, []int{1, 1}, 1, 3)(false, []int{})
}

func tCheckLine(t *testing.T, l []int, want, more int) func(bool, []int) {
	return func(b0 bool, l0 []int) {
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
