package tictactoe

import (
	"fmt"
	"strings"
)

func (g *Game) reverseOf(c Side) Side {
	if c == O {
		return X
	} else if c == X {
		return O
	}

	return Blank
}

func (g *Game) nextSide() Side {
	// Game ended/not started
	if len(g.History) == boardSize*boardSize || g.Status != StatusStarted {
		return Blank
	}

	// Game started, X first
	if len(g.History) == 0 {
		return X
	}

	lastSide := Side(g.History[len(g.History)-1][0])

	return g.reverseOf(lastSide)
}

func (g *Game) errorOf(num Code) error {
	return fmt.Errorf(g.errorDict()[num])
}

func (g *Game) nameOf(x, y int) string {
	return fmt.Sprintf("%s%d", string(rune('a'+x)), y+1)
}

// indexOf translate a3 into 0,2
func (g *Game) indexOf(name string) (int, int) {
	t := strings.ToLower(name)
	return int(rune(t[0]) - 'a'), int(rune(t[1]) - '1')
}

func (g *Game) sideOf(playerID string) Side {
	for i, p := range g.Players {
		if p == playerID {
			if i == 0 {
				return X
			}

			return O
		}
	}

	return Blank
}
