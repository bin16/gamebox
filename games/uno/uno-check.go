package uno

func checkNextCard(h history, card int) bool {
	pTurn := h[len(h)-1]
	pAction := pTurn[1]
	if pAction == ActionPlayCard {
		pCard := pTurn[2]
		pColor, pNum := cardColorAndNum(pCard)
		color, num := cardColorAndNum(card)
		if pCard == CardWild { // Wild x Normal
			pColor = pTurn[2]
		} else if isDrawTwo(pCard) { // DrawTwo x DrawTwo
			if !isDrawTwo(card) {
				return false
			}
		}

		// WildDrawFour x WildDrawFour
		// Normal x WildDrawFour
		if card == CardWild || card == CardWildDrawFour {
			return true
		}

		if pColor != color && pNum != num {
			return false
		}
	}

	return true
}
