package reversi

import (
	"encoding/json"
)

func (g *Game) Save() []byte {
	data, _ := json.Marshal(*g)

	return data
}

func Load(data []byte) (*Game, error) {
	g1 := &Game{}
	err := json.Unmarshal(data, g1)

	return g1, err
}
