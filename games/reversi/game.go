package reversi

import (
	"encoding/json"
)

func NewGame() *Game {
	return &Game{}
}

func Load(data []byte) (*Game, error) {
	m := map[string]interface{}{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	g := Game{
		board:   m["board"].(board),
		history: m["history"].(history),
		status:  m["status"].(int),
		players: m["players"].(map[string]int),
	}

	return &g, nil
}

type Game struct {
	board   board
	history history
	status  int
	players map[string]int
}

// Play : playerSide, boardIndex
func (g *Game) Play(commands []int) int {
	b, h, s, _ := g.flat()
	p := commands[0]
	n := commands[1]
	r := testStep(b, h, s, p, n)
	if r != OK {
		return r
	}

	b1, h1, s1, _ := commitStep(b, h, p, n)
	g.board = b1
	g.history = h1
	g.status = s1

	return OK
}

func (g *Game) Join(id string) int {
	pl := g.players
	if len(pl) == 2 {
		return NoChairs
	}
	if _, inGame := g.players[id]; inGame {
		return AlreadyIn
	}

	if len(pl) == 0 {
		g.players[id] = SideBlack
	} else if len(pl) == 1 {
		g.players[id] = SideWhite
	}

	return OK
}

func (g *Game) Status() int {
	return g.status
}

func (g *Game) Save() []byte {
	m := map[string]interface{}{
		"board":   g.board,
		"history": g.history,
		"status":  g.status,
		"players": g.players,
	}
	data, _ := json.Marshal(m)
	return data
}

func (g *Game) Data(id string) map[string]interface{} {
	m := map[string]interface{}{
		"board":   g.board,
		"history": g.history,
		"status":  g.status,
	}

	s, inGame := g.players[id]
	if !inGame {
		return m
	}

	b, h, _, _ := g.flat()
	ns := nextSide(b, h)
	m["side"] = s
	m["next"] = ns
	if s != Blank && ns == s {
		m["cells"] = getLiveCells(b, s)
	} else {
		m["cells"] = []int{}
	}

	return m
}

func (g *Game) flat() (board, history, int, map[string]int) {
	return g.board, g.history, g.status, g.players
}
