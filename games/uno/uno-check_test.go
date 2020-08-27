package uno

import (
	"testing"
)

func TestCheckNextCard(t *testing.T) {
	var h history

	// play_card normal
	h = history{
		[]int{0, ActionPlayCard, CardRedEight},
	}
	tCheckNextCard(t, h)(CardYellowEight, true)
	tCheckNextCard(t, h)(CardBlueEight, true)
	tCheckNextCard(t, h)(CardRedTwo, true)
	tCheckNextCard(t, h)(CardRedDrawTwo, true)
	tCheckNextCard(t, h)(CardRedReverse, true)
	tCheckNextCard(t, h)(CardRedSkip, true)
	tCheckNextCard(t, h)(CardWildDrawFour, true)
	tCheckNextCard(t, h)(CardWild, true)
	tCheckNextCard(t, h)(CardBlueTwo, false)
	tCheckNextCard(t, h)(CardYellowDrawTwo, false)
	tCheckNextCard(t, h)(CardGreenReverse, false)
	tCheckNextCard(t, h)(CardBlueSkip, false)

	// play_card yellow+draw_two
	h = history{
		[]int{0, ActionPlayCard, CardYellowDrawTwo},
	}
	tCheckNextCard(t, h)(CardYellowDrawTwo, true)
	tCheckNextCard(t, h)(CardYellowReverse, false)
	tCheckNextCard(t, h)(CardYellowZero, false)
	tCheckNextCard(t, h)(CardYellowNine, false)
	tCheckNextCard(t, h)(CardRedReverse, false)
	tCheckNextCard(t, h)(CardRedZero, false)
	tCheckNextCard(t, h)(CardRedNine, false)
	tCheckNextCard(t, h)(CardWild, false)
	tCheckNextCard(t, h)(CardWildDrawFour, false)

	// play_card wild+red,
	h = history{
		[]int{0, ActionPlayCard, CardWild, ColorRed},
	}
	tCheckNextCard(t, h)(CardWild, true)
	tCheckNextCard(t, h)(CardWildDrawFour, true)
	tCheckNextCard(t, h)(CardRedEight, true)
	tCheckNextCard(t, h)(CardRedSkip, true)
	tCheckNextCard(t, h)(CardYellowDrawTwo, false)
	tCheckNextCard(t, h)(CardYellowFour, false)
	tCheckNextCard(t, h)(CardGreenReverse, false)
	tCheckNextCard(t, h)(CardGreenSix, false)
	tCheckNextCard(t, h)(CardBlueDrawTwo, false)

	// play_card wild_draw_four+red
	h = history{
		[]int{0, ActionPlayCard, CardWildDrawFour, ColorBlue},
	}
	tCheckNextCard(t, h)(CardWildDrawFour, true)
	tCheckNextCard(t, h)(CardWild, false)
	tCheckNextCard(t, h)(CardRedEight, false)
	tCheckNextCard(t, h)(CardRedSkip, false)
	tCheckNextCard(t, h)(CardYellowDrawTwo, false)
	tCheckNextCard(t, h)(CardYellowFour, false)
	tCheckNextCard(t, h)(CardGreenReverse, false)
	tCheckNextCard(t, h)(CardGreenSix, false)
	tCheckNextCard(t, h)(CardBlueDrawTwo, false)
}

func tCheckNextCard(t *testing.T, h history) func(int, bool) {
	return func(nextCard int, ok bool) {
		t.Helper()
		if ok1 := checkNextCard(h, nextCard); ok1 != ok {
			t.Errorf("Failed: checkNextCard(h, %d), got %v, want %v;\n%d=%s", nextCard, ok1, ok, nextCard, nameOfCard(nextCard))
		}
	}
}
