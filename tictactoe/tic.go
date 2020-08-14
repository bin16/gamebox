package tictactoe

type Game struct {
	Status
	Board   [boardSize][boardSize]Side
	Players []string
	History [][3]int
}

func NewGame() *Game {
	return &Game{Status: StatusOpen}
}
