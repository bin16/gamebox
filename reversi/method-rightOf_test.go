package reversi

import "testing"

func TestRightOf(t *testing.T) {
	b := Black
	w := White
	o := Black
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

	x := 3
	y := 3
	if n := g.nameOf(3, 3); n != "d4" {
		t.Errorf("g.nameOf failed, %d,%d; %s", x, y, n)
	}
	if l := g.topOf(3, 3); !eqSideSlice(l, []Side{w, o, o}) {
		t.Errorf("g.topOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.topRightOf(3, 3); !eqSideSlice(l, []Side{b, w, b}) {
		t.Errorf("g.topRightOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.rightOf(3, 3); !eqSideSlice(l, []Side{w, b, w, b}) {
		t.Errorf("g.rightOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.bottomRightOf(3, 3); !eqSideSlice(l, []Side{b, b, w, b}) {
		t.Errorf("g.bottomRightOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.bottomOf(3, 3); !eqSideSlice(l, []Side{w, b, b, b}) {
		t.Errorf("g.bottomOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.bottomLeftOf(3, 3); !eqSideSlice(l, []Side{w, w, w}) {
		t.Errorf("g.bottomLeftOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.leftOf(3, 3); !eqSideSlice(l, []Side{o, o, o}) {
		t.Errorf("g.leftOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.topLeftOf(3, 3); !eqSideSlice(l, []Side{b, o, o}) {
		t.Errorf("g.topLeftOf failed: %d,%d; %v", x, y, l)
	}

	x = 7
	y = 7
	if n := g.nameOf(x, y); n != "h8" {
		t.Errorf("g.nameOf failed, %d,%d; %s", x, y, n)
	}
	if l := g.topRightOf(x, y); !eqSideSlice(l, []Side{}) {
		t.Errorf("g.topRightOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.topLeftOf(x, y); !eqSideSlice(l, []Side{w, b, b, b, b, o, o}) {
		t.Errorf("g.topLeftOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.leftOf(x, y); !eqSideSlice(l, []Side{w, w, w, b, b, b, b}) {
		t.Errorf("g.leftOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.rightOf(x, y); !eqSideSlice(l, []Side{}) {
		t.Errorf("g.rightOf failed: %d,%d; %v", x, y, l)
	}

	x = 7
	y = 0
	if n := g.nameOf(x, y); n != "h1" {
		t.Errorf("g.nameOf failed, %d,%d; %s", x, y, n)
	}
	if l := g.rightOf(x, y); !eqSideSlice(l, []Side{}) {
		t.Errorf("g.rightOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.leftOf(x, y); !eqSideSlice(l, []Side{b, b, b, o, b, o, o}) {
		t.Errorf("g.leftOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.bottomLeftOf(x, y); !eqSideSlice(l, []Side{w, w, w, w, o, b, b}) {
		t.Errorf("g.bottomLeftOf failed: %d,%d; %v", x, y, l)
	}
	if l := g.topOf(x, y); !eqSideSlice(l, []Side{}) {
		t.Errorf("g.topOf failed: %d,%d; %v", x, y, l)
	}
}

func eqSideSlice(a, b []Side) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
