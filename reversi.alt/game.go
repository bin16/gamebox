package reversi

type side int
type board [boardSize * boardSize]int
type history [][3]int

// Constants
const (
	Blank = iota
	SideBlack
	SideWhite

	boardSize   = 8
	top         = -boardSize
	topRight    = 1 - boardSize
	right       = 1
	bottomRight = boardSize + 1
	bottom      = boardSize
	bottomLeft  = boardSize - 1
	left        = -1
	topLeft     = -boardSize - 1

	StatusOpen
	StatusStarted
	StatusEnd

	OK
	NoChairs
	AlreadyIn
	NotStarted
	NotYourTurn
	NotFreeCell
)
