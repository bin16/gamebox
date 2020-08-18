package reversi

// Join check and add a playerID into game, and
// automatically change game status to GameStatusStarted
func (g *Game) Join(playerID string) error {
	if g.Status != GameStatusOpen {
		return g.errorOf(ErrorGameNotOpen)
	}

	for _, p := range g.Players {
		if p == playerID {
			return g.errorOf(ErrorPlayerAlreadyIn)
		}
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

func (g *Game) Data(playerID string) map[string]interface{} {
	data := map[string]interface{}{
		"id":     g.ID,
		"status": g.Status,
		"board":  g.Board,
		"next":   g.nextSide(),

		"side":    Blank,
		"cells":   [][2]int{},
		"options": []string{},
		"names":   g.Players,
		"scores":  g.scoreOf(),
		"history": g.History,
		"raw":     g,
	}
	if !g.playerInGame(playerID) {
		return data
	}

	c := g.sideOfPlayer(playerID)
	data["side"] = c
	if c != g.nextSide() {
		return data
	}

	data["cells"] = g.getAccessibleCells(c)
	data["options"] = g.getOptions(c)
	return data
}
