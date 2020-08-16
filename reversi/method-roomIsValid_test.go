package reversi

import "testing"

func TestRoomIsValid(t *testing.T) {
	b := Black
	w := White
	o := Blank
	g := NewGame()
	g.Board = [8][8]Side{
		[8]Side{o, o, o, o, b, o, w, b},
		[8]Side{o, o, o, o, o, w, b, b},
		[8]Side{b, b, b, o, w, o, b, b},
		[8]Side{o, o, w, b, w, b, b, b},
		[8]Side{b, b, b, w, b, b, b, w},
		[8]Side{b, w, w, b, b, b, w, w},
		[8]Side{b, w, w, w, w, w, w, w},
		[8]Side{b, b, b, b, b, b, b, b},
	}
	g.History = [][3]int{
		[3]int{int(w), 0, 0},
	}

	x := 3
	y := 1
	if !g.roomIsValid(b, x, y) {
		t.Errorf("g.roomIsValid failed: %s", g.nameOf(x, y))
	}

	x = 1
	y = 3
	if !g.roomIsValid(b, x, y) {
		t.Errorf("g.roomIsValid failed: %s", g.nameOf(x, y))
	}

	x = 2
	y = 3
	if !g.roomIsValid(b, x, y) {
		t.Errorf("g.roomIsValid failed: %s", g.nameOf(x, y))
	}

	x = 1
	y = 4
	if !g.roomIsValid(b, x, y) {
		t.Errorf("g.roomIsValid failed: %s", g.nameOf(x, y))
	}

	x = 0
	y = 5
	if !g.roomIsValid(b, x, y) {
		t.Errorf("g.roomIsValid failed: %s", g.nameOf(x, y))
	}

	x = 2
	y = 5
	if !g.roomIsValid(b, x, y) {
		t.Errorf("g.roomIsValid failed: %s", g.nameOf(x, y))
	}

	x = 0
	y = 0
	if g.roomIsValid(b, x, y) {
		t.Errorf("g.roomIsValid failed: %s", g.nameOf(x, y))
	}

	x = 7
	y = 7
	if g.roomIsValid(b, x, y) {
		t.Errorf("g.roomIsValid failed: %s", g.nameOf(x, y))
	}

	x = 1
	y = 5
	if g.roomIsValid(b, x, y) {
		t.Errorf("g.roomIsValid failed: %s", g.nameOf(x, y))
	}

	if l := g.getAccessibleCells(b); len(l) != 6 {
		t.Errorf("g.getAccessibleCells failed: %v", l)
	}
}
