package reversi

import (
	"fmt"
)

func (g *Game) print() {
	last := g.History[len(g.History)-1]
	cn := g.nextPlayer()
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
%s================`, score[0], score[1], pieceFace(g.nextPlayer()), s, history)
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
