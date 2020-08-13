package reversi

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

func (g *Game) Join(playerID string) error {
	if g.Status != GameStatusOpen {
		return fmt.Errorf("GAME_NOT_OPEN")
	}

	if len(g.Players) >= 2 {
		return fmt.Errorf("GAME_IS_FULL")
	}

	g.Players = append(g.Players, playerID)
	if len(g.Players) == 2 {
		g.Status = GameStatusStarted
	}

	return nil
}

func (g *Game) SideOfPlayer(playerID string) Side {
	if !g.playerInGame(playerID) {
		return Blank
	}

	for i, id := range g.Players {
		if id == playerID {
			return Side(i + 1)
		}
	}

	return Blank
}

func (g *Game) InfoOfPlayer(playerID string) map[string]interface{} {
	if !g.playerInGame(playerID) {
		return map[string]interface{}{
			"id":     playerID,
			"side":   Blank,
			"open":   false,
			"roomes": [][2]int{},
		}
	}

	c := g.SideOfPlayer(playerID)
	if c != g.nextPlayer() {
		return map[string]interface{}{
			"id":    playerID,
			"side":  int(c),
			"open":  false,
			"rooms": [][2]int{},
		}
	}

	return map[string]interface{}{
		"id":    playerID,
		"side":  int(c),
		"open":  true,
		"rooms": g.getRooms(c),
	}
}
