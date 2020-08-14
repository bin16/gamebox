package tictactoe

func (g *Game) playerInGame(playerID string) bool {
	for _, p := range g.Players {
		if p == playerID {
			return true
		}
	}

	return false
}

func (g *Game) Play(playerID, cmd string) error {
	if g.Status != StatusStarted {
		return g.errorOf(ErrorGameNotStarted)
	}

	if g.nextSide() != g.sideOf(playerID) {
		return g.errorOf(ErrorNotYourTurn) // Not your turn
	}

	x, y := g.indexOf(cmd)
	if x < 0 || x > boardSize-1 || y < 0 || y > boardSize-1 {
		return g.errorOf(ErrorBadCommand) // Bad cmd
	}

	s := g.sideOf(playerID)
	g.Board[x][y] = s
	g.History = append(g.History, [3]int{int(s), x, y})

	if len(g.History) == boardSize*boardSize {
		g.Status = StatusEnded
		return g.errorOf(GameEnd) // TODO: check here
	}

	return nil
}

func (g *Game) Join(playerID string) error {
	if g.Status != StatusOpen {
		return g.errorOf(ErrorGameNotOpen)
	}
	if len(g.Players) >= 2 {
		return g.errorOf(ErrorGameIsFull)
	}
	if g.playerInGame(playerID) {
		return g.errorOf(ErrorPlayerExists)
	}
	g.Players = append(g.Players, playerID)
	if len(g.Players) == 2 {
		g.Status = StatusStarted
	}

	return nil
}
