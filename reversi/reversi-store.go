package reversi

import (
	"encoding/json"
)

// Save export game data as []byte
func (g *Game) Save() []byte {
	data, _ := json.Marshal(*g)

	return data
}

// Load parse []byte data and open a game instance
func Load(data []byte) (*Game, error) {
	g1 := &Game{}
	err := json.Unmarshal(data, g1)

	return g1, err
}
