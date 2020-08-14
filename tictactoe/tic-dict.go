package tictactoe

// Side of Game
type Side int

// Status of Game
type Status int

// Code of errors
type Code int

// Constants
const (
	// Sides
	Blank Side = iota
	X
	O
	// Status
	StatusOpen Status = iota
	StatusStarted
	StatusEnded
	// Errors
	ErrorGameNotOpen    Code = iota // Failed to join
	ErrorPlayerExists               // Failed to join
	ErrorGameIsFull                 // Failed to join
	ErrorGameNotStarted             // Failed to play
	ErrorNotYourTurn                // Failed to play
	ErrorBadCommand                 // Failed to play
	GameEnd

	boardSize = 3
)

type DictEntry map[int]string

func (g *Game) errorDict() map[Code]string {
	return map[Code]string{
		ErrorGameNotOpen:    "Game not open",
		ErrorGameIsFull:     "Game is full",
		ErrorPlayerExists:   "Player exists",
		ErrorGameNotStarted: "Game not started",
		ErrorNotYourTurn:    "It's not your turn",
		ErrorBadCommand:     "Bad Command",
	}
}

func (g *Game) statusDict() map[Status]string {
	return map[Status]string{
		StatusOpen:    "open",
		StatusStarted: "started",
		StatusEnded:   "ended",
	}
}

func (g *Game) sideDict() map[Side]string {
	return map[Side]string{
		Blank: "Blank",
		X:     "X",
		O:     "O",
	}
}
