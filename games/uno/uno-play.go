package uno

func drawCards(h history) int {
	pAction, pCard, _ := readHistory(h, len(h)-1)
	if pAction == ActionPlayCard {
		if isDrawTwo(pCard) {
			return -1 // TODO:
		} else if pCard == CardWildDrawFour {
			return -1
		}
	}

	// Challenge is auto

	return 2
}
