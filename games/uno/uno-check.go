package uno

func checkNextCard(h history, card int) bool {
	pAction, pCard, pSetColor := readHistory(h, len(h)-1)
	if pAction == ActionPlayCard {
		pColor, pNum := cardColorAndNum(pCard)
		color, num := cardColorAndNum(card)
		// DrawTwo x DrawTwo
		if isDrawTwo(pCard) {
			if isDrawTwo(card) {
				return true
			}
			return false
		}

		// WildDrawFour x WildDrawFour
		if pCard == CardWildDrawFour {
			if card == CardWildDrawFour {
				return true
			}
			return false
		}

		// Wild x WildDrawFour
		// Wild x Wild
		// Normal x WildDrawFour
		// Normal x Wild
		if card == CardWildDrawFour || card == CardWild {
			return true
		}

		if isWild(pCard) {
			pColor = pSetColor
		}

		// Normal x Normal
		if pColor != color && pNum != num {
			return false
		}
	}

	return true
}

func checkAction(h history, action int) bool {
	if action == ActionDraw {
		return true
	}

	pAction, pCard, _ := readHistory(h, len(h)-1)
	if pAction == ActionPlayCard {
		if action == ActionChallenge {
			return pCard == CardWildDrawFour
		}

		if action == ActionCheckUNO || action == ActionPlayCard {
			return true
		}
	}

	return false
}
