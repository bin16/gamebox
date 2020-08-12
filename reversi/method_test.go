package reversi

import (
	"fmt"
)

func (g *Game) print() {
	cn := g.nextPlayer()
	var cw, cb int
	s := ""
	for y := 0; y < 8; y++ {
		r := fmt.Sprintf("|%d|", y+1)
		for x := 0; x < 8; x++ {
			c := g.Board[x][y]
			if c == White {
				cw++
				r += "w|"
			} else if c == Black {
				cb++
				r += "b|"
			} else if cn != Blank && g.roomIsValid(cn, x, y) {
				r += "O|"
			} else {
				r += "_|"
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

	fmt.Printf(`================
w: %d, b %d
|_|a_b_c_d_e_f_g_h|_|
%s|_|0_1_2_3_4_5_6_7|_|
----
%s================`, cw, cb, s, history)
}
