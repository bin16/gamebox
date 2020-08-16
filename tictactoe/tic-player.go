package tictactoe

import (
	"fmt"
)

func (g *Game) playerInGame(playerID string) bool {
	for _, p := range g.Players {
		if p == playerID {
			return true
		}
	}

	return false
}

func (g *Game) winnerSide() Side {
	flat4int := func(d [4]int) (int, int, int, int) {
		return d[0], d[1], d[2], d[3]
	}
	dl := [][4]int{
		[4]int{0, 0, 1, 0},
		[4]int{0, 0, 1, 1},
		[4]int{0, 0, 0, 1},
		[4]int{0, 1, 1, 0},
		[4]int{0, 2, 1, 0},
		[4]int{0, 2, 1, -1},
	}
	for _, d := range dl {
		x0, y0, dx, dy := flat4int(d)
		c0 := g.Board[x0][y0]
		l := g.lineOf(x0, y0, dx, dy, boardSize)
		for i, c := range l {
			if c != c0 {
				break
			}
			if c == c0 && i == len(l)-1 {
				return c0
			}
		}
	}
	fmt.Println("XXXXXXX", "end", dl, g.Board)

	return Blank
}

func (g *Game) lineOf(x0, y0, dx, dy, max int) []Side {
	l := []Side{}
	for i := 0; i < max; i++ {
		x := x0 + dx*i
		y := y0 + dy*i
		if len(l) == max || x == boardSize || y == boardSize {
			return l
		}
		c := g.Board[x][y]
		l = append(l, c)
	}

	return l
}

func (g *Game) isWin(x, y int) bool {
	dl := [][4]int{
		[4]int{0, y, 1, 0}, // right
		[4]int{x, 0, 0, 1}, // bottom
	}
	if y == x {
		dl = append(dl, [4]int{0, 0, 1, 1}) // bottomRight
	} else if x+y == boardSize-1 {
		dl = append(dl, [4]int{0, 2, 1, -1}) // topRight
	}

	c0 := g.Board[x][y]
	for _, p := range dl {
		l := g.lineOf(p[0], p[1], p[2], p[3], boardSize)
		for i, c := range l {
			if c != c0 {
				return false
			}
			if c == c0 && i == len(l)-1 {
				return true
			}
		}
	}

	return true
}

func (g *Game) Play(playerID, cmd string) error {
	if g.Status != StatusStarted {
		return g.errorOf(ErrorGameNotStarted)
	}

	if g.nextSide() != g.sideOf(playerID) {
		return g.errorOf(ErrorNotYourTurn) // Not your turn
	}

	x, y := g.indexOf(cmd)
	if x < 0 || x > boardSize-1 || y < 0 || y > boardSize-1 {
		return g.errorOf(ErrorBadCommand) // Bad cmd
	}

	s := g.sideOf(playerID)
	g.Board[x][y] = s
	g.History = append(g.History, [3]int{int(s), x, y})

	if len(g.History) == boardSize*boardSize || g.isWin(x, y) {
		g.Status = StatusEnded
	}

	return nil
}

func (g *Game) Join(playerID string) error {
	if g.Status != StatusOpen {
		return g.errorOf(ErrorGameNotOpen)
	}
	if len(g.Players) >= 2 {
		return g.errorOf(ErrorGameIsFull)
	}
	if g.playerInGame(playerID) {
		return g.errorOf(ErrorPlayerExists)
	}
	g.Players = append(g.Players, playerID)
	if len(g.Players) == 2 {
		g.Status = StatusStarted
	}

	return nil
}

func (g *Game) End() (bool, Side) {
	return g.Status == StatusEnded, g.winnerSide()
}

func (g *Game) Info(playerID string) map[string]interface{} {
	return map[string]interface{}{
		"board":  g.Board,
		"status": g.Status,
		"player": map[string]interface{}{
			"id":   playerID,
			"open": g.nextSide() == g.sideOf(playerID),
		},
	}
}
