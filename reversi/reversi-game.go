package reversi

import (
	"github.com/google/uuid"
)

// NewGame -> New Reversi Game
func NewGame() *Game {
	id := uuid.Must(uuid.NewRandom())
	g := &Game{
		ID:     id.String(),
		Status: GameStatusOpen,
	}
	g.Board[3][3] = Black
	g.Board[4][4] = Black
	g.Board[3][4] = White
	g.Board[4][3] = White
	g.History = [][3]int{
		[3]int{int(Black), 3, 3},
		[3]int{int(White), 3, 4},
		[3]int{int(Black), 4, 4},
		[3]int{int(White), 4, 3},
	}
	g.Players = []string{}

	return g
}

// Game of Reversi
type Game struct {
	ID      string     `json:"id"`
	Status  Status     `json:"status"`
	Board   [8][8]Side `json:"board"`
	History [][3]int   `json:"history"` // Color,X,Y
	Players []string   `json:"players"`
}

// getRooms return all available cells, for specific color
func (g *Game) getRooms(c Side) [][2]int {
	cells := [][2]int{}
	for x := 0; x < len(g.Board); x++ {
		for y := 0; y < len(g.Board[0]); y++ {
			if g.roomIsValid(c, x, y) {
				cells = append(cells, [2]int{x, y})
			}
		}
	}

	return cells
}

func (g *Game) lineIsValid(c Side, line []Side) bool {
	if len(line) < 2 {
		return false
	}

	for i, cn := range line {
		if cn == Blank {
			return false
		} else if cn == c {
			return i > 0
		}
	}

	return false
}

// roomIsValid check if we can put Color C into Board x,y
func (g *Game) roomIsValid(c Side, x, y int) bool {
	if g.Board[x][y] != Blank {
		return false
	}

	if g.lineIsValid(c, g.topOf(x, y)) {
		return true
	}

	if g.lineIsValid(c, g.topRightOf(x, y)) {
		return true
	}

	if g.lineIsValid(c, g.rightOf(x, y)) {
		return true
	}

	if g.lineIsValid(c, g.bottomRightOf(x, y)) {
		return true
	}

	if g.lineIsValid(c, g.bottomOf(x, y)) {
		return true
	}

	if g.lineIsValid(c, g.bottomLeftOf(x, y)) {
		return true
	}

	if g.lineIsValid(c, g.leftOf(x, y)) {
		return true
	}

	if g.lineIsValid(c, g.topLeftOf(x, y)) {
		return true
	}

	return false
}

func (g *Game) reverseOf(c Side) Side {
	if c == White {
		return Black
	}

	return White
}

func (g *Game) topOf(x, y int) []Side {
	line, _ := g.lineOf(x, y, 0, -1)

	return line
}

func (g *Game) topRightOf(x, y int) []Side {
	line, _ := g.lineOf(x, y, 1, -1)

	return line
}

func (g *Game) rightOf(x, y int) []Side {
	line, _ := g.lineOf(x, y, 1, 0)

	return line
}

func (g *Game) bottomRightOf(x, y int) []Side {
	line, _ := g.lineOf(x, y, 1, 1)

	return line
}

func (g *Game) bottomOf(x, y int) []Side {
	line, _ := g.lineOf(x, y, 0, 1)

	return line
}

func (g *Game) bottomLeftOf(x, y int) []Side {
	line, _ := g.lineOf(x, y, -1, 1)

	return line
}

func (g *Game) leftOf(x, y int) []Side {
	line, _ := g.lineOf(x, y, -1, 0)

	return line
}

func (g *Game) topLeftOf(x, y int) []Side {
	line, _ := g.lineOf(x, y, -1, -1)

	return line
}

func (g *Game) nextPlayer() Side {
	if g.Status != GameStatusStarted {
		return Blank
	}

	p1c := g.History[len(g.History)-1][0]
	p1 := Side(p1c)

	p2 := g.reverseOf(p1)
	if len(g.getRooms(p2)) > 0 {
		return p2
	} else if len(g.getRooms(p1)) > 0 {
		return p1
	}

	return Blank // Game Ended
}

func (g *Game) lineOf(x, y int, dx, dy int) ([]Side, [][2]int) {
	pieces := []Side{}
	line := [][2]int{}

	if dx == 0 && dy == 0 {
		return pieces, line
	}

	x1 := x
	y1 := y
	for {
		x1 += dx
		y1 += dy
		if x1 < 0 || y1 < 0 || x1 > 7 || y1 > 7 {
			break
		}
		pieces = append(pieces, g.Board[x1][y1])
		line = append(line, [2]int{x1, y1})
	}

	return pieces, line
}
