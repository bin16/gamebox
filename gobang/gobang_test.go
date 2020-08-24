package tic

import (
	"testing"
)

func testBoard() (board, history) {
	W := SideWhite
	B := SideBlack
	r := Blank
	var b board = [boardSize * boardSize]int{
		W, B, W, W, B, B, W, r, r, r, r, r, r, r, r, r, r, r, r,
		r, B, r, W, r, B, B, r, r, r, r, r, r, r, r, r, r, r, r,
		W, B, B, B, W, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		B, r, W, r, B, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, B, r, W, B, B, W, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, B, r, W, r, B, B, r, r, r, r, r, r, r, r,
		r, r, r, r, r, B, B, B, W, r, B, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
		r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r, r,
	}
	var h history = [][3]int{
		[3]int{SideWhite, 0, 0}, // O
	}

	return b, h
}

func TestMethodCommitStep(t *testing.T) {
	b, h := testBoard()
	tCommitStep(t, b, h, SideBlack, 19*19-1)(StatusStarted, SideWhite)
	tCommitStep(t, b, h, SideWhite, 19*19-1)(StatusStarted, SideBlack)
	tCommitStep(t, b, h, SideBlack, 19*5+6)(StatusEnd, SideBlack)
	tCommitStep(t, b, h, SideWhite, 19*3+5)(StatusEnd, SideWhite)
	tCommitStep(t, b, h, SideWhite, 19*3+1)(StatusStarted, SideBlack)
	tCommitStep(t, b, h, SideBlack, 19*3+1)(StatusEnd, SideBlack)
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
	b, _ := testBoard()
	r := Blank
	B := SideBlack
	W := SideWhite
	tPickLine(t, b, 19*4+7, bottomRight)([]int{r, r, r, r, r, r, r, r, r, r, r, r})
	tPickLine(t, b, 19*1+3, bottom)([]int{W, B, r, W, r, r, r, r, r, r, r, r, r, r, r, r, r, r})
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
		SideWhite, Blank, SideWhite, Blank, SideBlack, Blank, Blank, Blank, SideBlack,
	}
	var h history = [][3]int{
		[3]int{SideBlack, 4, 0}, // X
		[3]int{SideWhite, 2, 0}, // O
		[3]int{SideBlack, 8, 0}, // X
		[3]int{SideWhite, 0, 0}, // O
	}
	var s = StatusStarted

	if nextSide(h) != SideBlack {
		t.Errorf("nextSide(h) should be %d, got %d", SideBlack, nextSide(h))
	}

	tTestStep(t, b, h, StatusOpen, SideBlack, 7)(NotStarted)
	tTestStep(t, b, h, StatusEnd, SideBlack, 7)(NotStarted)
	tTestStep(t, b, h, s, SideBlack, 7)(OK)
	tTestStep(t, b, h, s, SideBlack, 6)(OK)
	tTestStep(t, b, h, s, SideBlack, 5)(OK)
	tTestStep(t, b, h, s, SideBlack, 4)(NotFreeCell)
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
	b, h := testBoard()
	if s, ws := checkBoard(b, SideBlack); s != StatusStarted || ws != Blank {
		t.Errorf("Failed: checkBoard(b, %d), got (%d, %d), want (%d, %d)", SideBlack, s, ws, StatusStarted, Blank)
	}
	if s, ws := checkBoard(b, SideWhite); s != StatusStarted || ws != Blank {
		t.Errorf("Failed: checkBoard(b, %d), got (%d, %d), want (%d, %d)", SideWhite, s, ws, StatusStarted, Blank)
	}

	b1, _, _, _ := commitStep(b, h, SideWhite, 19*3+5)
	if s, ws := checkBoard(b1, SideWhite); s != StatusEnd || ws != SideWhite {
		t.Errorf("Failed: checkBoard(b, %d), got (%d, %d), want (%d, %d)", SideWhite, s, ws, StatusEnd, SideWhite)
	}
}

func TestMethodCheckLine(t *testing.T) {
	noResult := [2]int{}
	tCheckLine(t, []int{2, 2, 0, 2, 1, 1, 2, 2, 0, 2, 2, 1, 2, 1, 2}, 2, lenOfWin)(false, noResult)
	tCheckLine(t, []int{2, 2, 0, 2, 1, 1, 2, 2, 2, 2, 2, 1, 2, 1, 2}, 2, lenOfWin)(true, [2]int{6, 11})
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
