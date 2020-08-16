package reversi

// Side of Game
type Side int

// Status of Game
type Status int

// ErrorCode of Game
type ErrorCode int

// Game Codes
const (
	// Sides
	Blank Side = iota
	Black
	White
	// Status
	GameStatusOpen Status = iota
	GameStatusReady
	GameStatusStarted
	GameStatusEnded
	// Error Codes
	ErrorPlayerAlreadyIn ErrorCode = iota
	ErrorNotInGame
	ErrorNotYourTurn
	ErrorBadRequest
	ErrorGameNotOpen
	ErrorGameIsFull
	ErrorGameEnded
)

// DictEntry has a key of int codes and value of string
type DictEntry map[interface{}]string

// Dict return translations of constants
func (g *Game) Dict() DictEntry {
	return DictEntry{
		Blank:                "blank",
		Black:                "black",
		White:                "white",
		GameStatusOpen:       "open",
		GameStatusReady:      "ready",
		GameStatusStarted:    "started",
		GameStatusEnded:      "ended",
		ErrorPlayerAlreadyIn: "Player already in",
		ErrorNotInGame:       "You are not in game",
		ErrorNotYourTurn:     "It's not your turn",
		ErrorBadRequest:      "Bad Request",
		ErrorGameNotOpen:     "Failed to join game, it is not open",
		ErrorGameIsFull:      "Failed to join game, it is full",
		ErrorGameEnded:       "Game ended",
	}
}
