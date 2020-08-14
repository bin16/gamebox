package reversi

// Join check and add a playerID into game, and
// automatically change game status to GameStatusStarted
func (g *Game) Join(playerID string) error {
	if g.Status != GameStatusOpen {
		return g.errorOf(ErrorGameNotOpen)
	}

	g.Players = append(g.Players, playerID)
	if len(g.Players) == 2 {
		g.Status = GameStatusStarted
	}

	return nil
}

// TODO: remove
func (g *Game) playerInGame(playerID string) bool {
	for _, p := range g.Players {
		if p == playerID {
			return true
		}
	}

	return false
}

// sideOfPlayer return Side of player, Black of White,
// and Blank if player not in game
func (g *Game) sideOfPlayer(playerID string) Side {
	for i, id := range g.Players {
		if id == playerID {
			if i == 0 {
				return Black
			} else if i == 1 {
				return White
			}
		}
	}

	return Blank
}

// TODO: add dict
func (g *Game) InfoOfPlayer(playerID string) map[string]interface{} {
	if !g.playerInGame(playerID) {
		return map[string]interface{}{
			"id":     playerID,
			"side":   Blank,
			"open":   false,
			"roomes": [][2]int{},
		}
	}

	c := g.sideOfPlayer(playerID)
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
