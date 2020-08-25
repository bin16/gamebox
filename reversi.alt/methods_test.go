package reversi

import "testing"

func TestGetLiveCells(t *testing.T) {
	b, _ := testBoard()
	B := SideBlack
	W := SideWhite
	O := Blank

	tCellIsLive(t, b, boardSize*2+2, B)(false)
	tCellIsLive(t, b, boardSize*2+2, W)(false)
	tCellIsLive(t, b, boardSize*3+4, B)(false)
	tCellIsLive(t, b, boardSize*3+4, W)(false)
	tCellIsLive(t, b, 0, B)(false)
	tCellIsLive(t, b, 0, W)(true)
	// o, b, b, w, b, o, o, o
	tCheckLine(t, append([]int{B}, pickLine(b, 0, bottomRight)[1:]...))(false)
	tCheckLine(t, append([]int{W}, pickLine(b, 0, bottomRight)[1:]...))(true)
	tCheckLine(t, append([]int{B}, pickLine(b, 0, right)[1:]...))(false)
	tCheckLine(t, append([]int{W}, pickLine(b, 0, right)[1:]...))(false)
	tCheckLine(t, append([]int{B}, pickLine(b, 0, topRight)[1:]...))(false)
	tCheckLine(t, append([]int{W}, pickLine(b, 0, topRight)[1:]...))(false)
	tCheckLine(t, append([]int{B}, pickLine(b, 0, top)[1:]...))(false)
	tCheckLine(t, append([]int{W}, pickLine(b, 0, top)[1:]...))(false)
	tCheckLine(t, append([]int{B}, pickLine(b, 0, topLeft)[1:]...))(false)
	tCheckLine(t, append([]int{W}, pickLine(b, 0, topLeft)[1:]...))(false)
	tCheckLine(t, append([]int{B}, pickLine(b, 0, left)[1:]...))(false)
	tCheckLine(t, append([]int{W}, pickLine(b, 0, left)[1:]...))(false)
	tCheckLine(t, append([]int{B}, pickLine(b, 0, bottomLeft)[1:]...))(false)
	tCheckLine(t, append([]int{W}, pickLine(b, 0, bottomLeft)[1:]...))(false)
	tCheckLine(t, append([]int{B}, pickLine(b, 0, bottom)[1:]...))(false)
	tCheckLine(t, append([]int{W}, pickLine(b, 0, bottom)[1:]...))(false)
	tGetLiveCells(t, b, B)([]int{20, 22, 34, 44, 50})

	tCheckLine(t, []int{B, W, B})(true)
	tCheckLine(t, []int{W, B, W})(true)
	tCheckLine(t, []int{B, B, W})(false)
	tCheckLine(t, []int{W, B, B})(false)
	tCheckLine(t, []int{W, W, W})(false)
	tCheckLine(t, []int{W, B, O})(false)
	tCheckLine(t, []int{W, B, O})(false)
	tCheckLine(t, []int{O, W, W})(false)
	tCheckLine(t, []int{W, W, O})(false)
}

func tCheckLine(t *testing.T, l []int) func(r0 bool) {
	return func(r0 bool) {
		t.Helper()
		if r1 := checkLine(l); r1 != r0 {
			t.Errorf("Failed: checkLine(%v), got %v, want %v", l, r1, r0)
		}
	}
}

func tCellIsLive(t *testing.T, b board, n, s0 int) func(r0 bool) {
	return func(r0 bool) {
		t.Helper()
		if r1 := cellIsLive(b, n, s0); r1 != r0 {
			t.Errorf("Failed: cellIsLive(b, %d, %d), got %v, want %v", n, s0, r1, r0)
		}
	}
}

func tGetLiveCells(t *testing.T, b board, s0 int) func(l10 []int) {
	return func(l0 []int) {
		t.Helper()
		if l1 := getLiveCells(b, s0); len(l1) != len(l0) {
			t.Errorf("Failed: getLiveCells(b, %d), got %v, want %v", s0, l1, l0)
		} else {
			for i, c1 := range l1 {
				if c1 != l0[i] {
					t.Errorf("Failed: getLiveCells(b, %d), got %v, want %v", s0, l1, l0)
					return
				}
			}
		}
	}
}

func TestPickLine(t *testing.T) {
	b, _ := testBoard()
	o := Blank
	B := SideBlack
	W := SideWhite
	tPickLine(t, b, boardSize*3+3, top)([]int{W, W, o, o})
	tPickLine(t, b, boardSize*3+3, topRight)([]int{W, o, o, o})
	tPickLine(t, b, boardSize*3+3, right)([]int{W, W, W, W, o})
	tPickLine(t, b, boardSize*3+3, bottomRight)([]int{W, B, o, o, o})
	tPickLine(t, b, boardSize*3+3, bottom)([]int{W, W, W, W, o})
	tPickLine(t, b, boardSize*3+3, bottomLeft)([]int{W, o, o, o})
	tPickLine(t, b, boardSize*3+3, left)([]int{W, W, o, o})
	tPickLine(t, b, boardSize*3+3, topLeft)([]int{W, B, B, o})
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

func testBoard() (board, history) {
	b := SideBlack
	w := SideWhite
	o := Blank
	return board{
			o, o, o, o, o, o, o, o,
			o, b, o, o, o, o, o, o,
			b, b, b, w, o, o, o, o,
			o, o, w, w, w, w, w, o,
			o, w, o, w, b, o, o, o,
			o, o, w, w, o, o, o, o,
			o, o, o, w, o, o, o, o,
			o, o, o, o, o, o, o, o,
		}, history{
			[3]int{w, boardSize*4 + 1},
		}
}
