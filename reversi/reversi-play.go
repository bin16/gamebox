package reversi

import (
	"fmt"
	"strings"
)

func (g *Game) errorOf(errCode ErrorCode) error {
	return fmt.Errorf(g.Dict()[errCode])
}

// Play reversi TODO: comment
func (g *Game) Play(playerID, cmd string) error {
	c := g.sideOfPlayer(playerID)
	if c == Blank {
		return g.errorOf(ErrorNotInGame)
	}

	x, y := g.indexOf(cmd)
	if err := g.checkNext(c, x, y); err != nil {
		return err
	}

	cells := g.cellsToReserve(c, x, y)
	for _, p := range cells {
		g.Board[p[0]][p[1]] = c
	}
	g.Board[x][y] = c
	g.History = append(g.History, [3]int{int(c), x, y})

	return nil
}

// checkNext check if a side could be played
func (g *Game) checkNext(c Side, x, y int) error {
	n := g.nextPlayer()
	if n == Blank {
		return g.errorOf(ErrorGameEnded)
	}

	if n != c {
		return g.errorOf(ErrorNotYourTurn)
	}

	if x < 0 || x > 7 || y < 0 || y > 7 {
		return g.errorOf(ErrorBadRequest)
	}

	if !g.roomIsValid(c, x, y) {
		return g.errorOf(ErrorBadRequest)
	}

	return nil
}

func (g *Game) cellsToReserve(c Side, x, y int) [][2]int {
	cells := [][2]int{}
	d := []int{-1, 0, 1}
	for _, dx := range d {
		for _, dy := range d {
			pieces, line := g.lineOf(x, y, dx, dy)
			for i, p := range pieces {
				if p == c {
					cells = append(cells, line[:i]...)
					break
				} else if p == Blank {
					break
				}
			}
		}
	}

	return cells
}

// scoreOf() return piece count of Black, White and total
func (g *Game) scoreOf() [3]int {
	cb := 0
	cw := 0
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			b := g.Board[x][y]
			if b == Black {
				cb++
			} else if b == White {
				cw++
			}
		}
	}

	return [3]int{cb, cw, cb + cw}
}

// nameOf translate 2,3 into c4
func (g *Game) nameOf(x, y int) string {
	return fmt.Sprintf("%s%d", string('a'+x), y+1)
}

// indexOf translate c4 into 2,3
func (g *Game) indexOf(name string) (int, int) {
	t := strings.ToLower(name)
	return int(rune(t[0]) - 'a'), int(rune(t[1]) - '1')
}
