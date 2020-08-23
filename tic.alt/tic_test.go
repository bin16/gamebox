package tic

import (
	"testing"
)

func TestMethodCommitStep(t *testing.T) {
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
	tCommitStep(t, b, h, SideX, 1)(StatusStarted, SideO)
	tCommitStep(t, b, h, SideX, 3)(StatusStarted, SideO)
	tCommitStep(t, b, h, SideX, 5)(StatusStarted, SideO)
	tCommitStep(t, b, h, SideX, 6)(StatusStarted, SideO)
	tCommitStep(t, b, h, SideX, 7)(StatusStarted, SideO)

	tCommitStep(t, b, h, SideO, 1)(StatusEnd, SideO)
	tCommitStep(t, b, h, SideX, 0)(StatusEnd, SideX)

	// O _ O
	// _ X _
	// X _ X
	b1, h1, _, _ := commitStep(b, h, SideX, 6)
	tCommitStep(t, b1, h1, SideO, 1)(StatusEnd, SideO)
	tCommitStep(t, b1, h1, SideO, 7)(StatusStarted, SideX)

	tCommitStep(t, b1, h1, SideO, 1)(StatusEnd, SideO)
	tCommitStep(t, b1, h1, SideX, 7)(StatusEnd, SideX)
}

func tCommitStep(t *testing.T, b board, h history, p, n int) func(s0, ns0 int) {
	return func(s0, ns0 int) {
		t.Helper()
		if _, _, s1, ns1 := commitStep(b, h, p, n); s1 != s0 || ns1 != ns0 {
			t.Errorf("Failed: commitStep(..., %d, %d), got (..., %d, %d), want (..., %d, %d)", p, n, s1, ns1, s0, ns0)
		}
	}
}

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

func TestMethodTestStep(t *testing.T) {
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
	var s = StatusStarted

	if nextSide(h) != SideX {
		t.Errorf("nextSide(h) should be %d, got %d", SideX, nextSide(h))
	}

	tTestStep(t, b, h, StatusOpen, SideX, 7)(NotStarted)
	tTestStep(t, b, h, StatusDraw, SideX, 7)(NotStarted)
	tTestStep(t, b, h, StatusEnd, SideX, 7)(NotStarted)
	tTestStep(t, b, h, s, SideX, 7)(OK)
	tTestStep(t, b, h, s, SideX, 6)(OK)
	tTestStep(t, b, h, s, SideX, 5)(OK)
	tTestStep(t, b, h, s, SideX, 4)(NotFreeCell)
}

func tTestStep(t *testing.T, b board, h history, s, p, n int) func(r0 int) {
	return func(r0 int) {
		t.Helper()
		r1 := testStep(b, h, s, p, n)
		if r1 != r0 {
			t.Errorf("Failed: testStep(%v, h, %d, %d, %d) got (%d); want (%d)", b, s, p, n, r1, r0)
		}
	}
}

func TestMethodCheckBoard(t *testing.T) {
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

func TestMethodCheckLine(t *testing.T) {
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
