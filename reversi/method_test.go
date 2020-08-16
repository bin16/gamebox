package reversi

import (
	"fmt"
	"testing"
)

func (g *Game) print() {
	last := g.History[len(g.History)-1]
	cn := g.nextSide()
	s := ""
	for y := 0; y < 8; y++ {
		r := fmt.Sprintf("|%d|", y+1)
		for x := 0; x < 8; x++ {
			c := g.Board[x][y]
			if last[1] == x && last[2] == y {
				r += pieceHistoryFace(c)
			} else if cn != Blank {
				if g.roomIsValid(cn, x, y) {
					r += pieceOkFace(cn)
				} else {
					r += pieceFace(c)
				}
			} else {
				r += pieceFace(c)
			}
		}
		r += fmt.Sprintf("%d|\n", y)
		s += r
	}

	history := ""
	for i, h := range g.History {
		c := Side(h[0])
		x := h[1]
		y := h[2]
		if c == White {
			history += fmt.Sprintf("%d. w - %s\n", i+1, g.nameOf(x, y))
		} else {
			history += fmt.Sprintf("%d. b - %s\n", i+1, g.nameOf(x, y))
		}
	}

	score := g.scoreOf()
	fmt.Printf(`================
w: %d, b %d, now: %s
|_|a_b_c_d_e_f_g_h|_|
%s|_|0_1_2_3_4_5_6_7|_|
----
%s================`, score[0], score[1], pieceFace(g.nextSide()), s, history)
}

func pieceFace(c Side) string {
	if c == Black {
		return "🌑"
	} else if c == White {
		return "🌕"
	}

	return "⬜"
}

func pieceOkFace(c Side) string {
	return "🌙"
}

func pieceHistoryFace(c Side) string {
	if c == Black {
		return "🌚"
	} else if c == White {
		return "🌝"
	}

	return ""
}

func mustJoin(t *testing.T, g *Game) func(string) {
	return func(p string) {
		t.Helper()
		if err := g.Join(p); err != nil {
			t.Errorf("Failed to run g.Join(%s); got error <%v>", p, err)
		}
	}
}

func mustPlay(t *testing.T, g *Game) func(string, string) {
	return func(p, c string) {
		t.Helper()
		if err := g.Play(p, c); err != nil {
			t.Errorf("Failed to run g.Play(%s, %s); got error <%v>", p, c, err)
		}
	}
}

func assertGameIsEnd(t *testing.T, g *Game) func(bool, Side) {
	return func(e bool, c Side) {
		t.Helper()
		if e1, c1 := g.End(); e1 != e || c1 != c {
			t.Errorf("Failed to run g.End(); got (%v, %d); want (%v, %d)", e1, c1, e, c)
		}
	}
}

func assertGameStatusIs(t *testing.T, g *Game) func(s Status) {
	return func(s Status) {
		t.Helper()
		if g.Status != s {
			t.Errorf("Failed to get g.Status; got %d; want %d", g.Status, s)
		}
	}
}

func assertScoreIs(t *testing.T, g *Game) func([3]int) {
	return func(s [3]int) {
		t.Helper()
		s1 := g.scoreOf()
		for i, v := range s1 {
			if s[i] != v {
				t.Errorf("Failed to run g.scoreOf(), got %v, want %v", s1, s)
			}
		}
	}
}

func assertIndexOf(t *testing.T, g *Game) func(n string, x, y int) {
	return func(n string, x, y int) {
		t.Helper()
		if x1, y1 := g.indexOf(n); x1 != x || y1 != y {
			t.Errorf("Failed to run g.indexOf(%s); got (%d, %d); want (%d, %d)", n, x1, y1, x, y)
		}
	}
}

func assertNameOf(t *testing.T, g *Game) func(x, y int, n string) {
	return func(x, y int, n string) {
		t.Helper()
		if n1 := g.nameOf(x, y); n1 != n {
			t.Errorf("Failed to run g.nameOf(%d, %d); got (%s); want (%s)", x, y, n1, n)
		}
	}
}
