package reversi

import (
	"fmt"
	"strings"
)

// Play reversi TODO: comment
func (g *Game) Play(c Side, name string) error {
	x, y := g.indexOf(name)
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

func (g *Game) checkNext(c Side, x, y int) error {
	n := g.nextPlayer()
	if n == Blank {
		return fmt.Errorf("GAME_ENDED")
	}

	if g.nextPlayer() != c {
		return fmt.Errorf("NOT_YOUR_TURN")
	}

	if x < 0 || x > 7 || y < 0 || y > 7 {
		return fmt.Errorf("BAD_REQUEST")
	}

	if !g.roomIsValid(c, x, y) {
		return fmt.Errorf("BAD_REQUEST")
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
				if p == c && i >= 1 {
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

func (g *Game) nameOf(x, y int) string {
	return fmt.Sprintf("%s%d", string('a'+x), y+1)
}

func (g *Game) indexOf(name string) (int, int) {
	t := strings.ToLower(name)
	return int(rune(t[0]) - 'a'), int(rune(t[1]) - '1')
}
