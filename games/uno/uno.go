package uno

// Constants
const (
	CardRedZero = iota
	CardRedOne
	CardRedTwo
	CardRedThree
	CardRedFour
	CardRedFive
	CardRedSix
	CardRedSeven
	CardRedEight
	CardRedNine
	CardRedSkip
	CardRedReverse
	CardRedDrawTwo

	CardYellowZero
	CardYellowOne
	CardYellowTwo
	CardYellowThree
	CardYellowFour
	CardYellowFive
	CardYellowSix
	CardYellowSeven
	CardYellowEight
	CardYellowNine
	CardYellowSkip
	CardYellowReverse
	CardYellowDrawTwo

	CardGreenZero
	CardGreenOne
	CardGreenTwo
	CardGreenThree
	CardGreenFour
	CardGreenFive
	CardGreenSix
	CardGreenSeven
	CardGreenEight
	CardGreenNine
	CardGreenSkip
	CardGreenReverse
	CardGreenDrawTwo

	CardBlueZero
	CardBlueOne
	CardBlueTwo
	CardBlueThree
	CardBlueFour
	CardBlueFive
	CardBlueSix
	CardBlueSeven
	CardBlueEight
	CardBlueNine
	CardBlueSkip
	CardBlueReverse
	CardBlueDrawTwo

	CardWild
	CardWildDrawFour

	ColorRed
	ColorYellow
	ColorGreen
	ColorBlue
	ColorSpecial
	NotColor

	ModeStandard
	ModeEndless
	ModeEndlessBoom

	StatusOpen
	StatusStarted
	StatusEnded

	ActionDraw     // draw, draw+draw_two draw+wild_draw_four, draw+12, draw+N,
	ActionPlayCard // red+2, wild+red, yellow+skip, wild_draw_four+blue
	ActionUNO      // uno
	ActionCall     // call
	ActionServerCommit
)

var cardLibrary = []int{}
var zeroCards, specialCards map[int]int

type history [][3]int // player index, action, action param, time

func init() {
	zeroCards = map[int]int{
		CardRedZero:    1,
		CardYellowZero: 1,
		CardGreenZero:  1,
		CardBlueZero:   1,
	}
	specialCards = map[int]int{
		CardWild:         4,
		CardWildDrawFour: 4,
	}

	for i := CardRedZero; i < CardWildDrawFour+1; i++ {
		for r := 0; r < cardCount(i); r++ {
			cardLibrary = append(cardLibrary, i)
		}
	}
}

func zeros() []int {
	return []int{
		CardRedZero, CardYellowZero,
		CardGreenZero, CardBlueZero,
	}
}

func specials() []int {
	return []int{
		CardWild,
		CardWildDrawFour,
	}
}

func isWild(n int) bool {
	return n == CardWild || n == CardWildDrawFour
}

func isNum(n int) bool {
	if isWild(n) || isDrawTwo(n) || isReverse(n) || isSkip(n) {
		return false
	}

	return true
}

func isZero(n int) bool {
	_, ok := zeroCards[n]
	return ok
}

func isDrawTwo(n int) bool {
	m := map[int]int{
		CardRedDrawTwo:    1,
		CardYellowDrawTwo: 1,
		CardGreenDrawTwo:  1,
		CardBlueDrawTwo:   1,
	}
	_, ok := m[n]

	return ok
}

func isSkip(n int) bool {
	m := map[int]int{
		CardRedSkip:    1,
		CardYellowSkip: 1,
		CardGreenSkip:  1,
		CardBlueSkip:   1,
	}
	_, ok := m[n]

	return ok
}

func isReverse(n int) bool {
	m := map[int]int{
		CardRedReverse:    1,
		CardYellowReverse: 1,
		CardGreenReverse:  1,
		CardBlueReverse:   1,
	}
	_, ok := m[n]

	return ok
}

func isSpecial(n int) bool {
	_, ok := specialCards[n]
	return ok
}

func cardCount(n int) int {
	if isSpecial(n) {
		return 4
	} else if isZero(n) {
		return 1
	}

	return 2
}

func cardColorAndNum(n int) (color, num int) {
	color = cardColor(n)
	switch color {
	case ColorRed:
		num = n - CardRedZero
	case ColorYellow:
		num = n - CardYellowZero
	case ColorGreen:
		num = n - CardGreenZero
	case ColorBlue:
		num = n - CardBlueZero
	case ColorSpecial:
		num = n
	}

	return color, num
}

func cardColor(n int) int {
	if n >= CardRedZero && n <= CardRedDrawTwo {
		return ColorRed
	}

	if n >= CardYellowZero && n <= CardYellowDrawTwo {
		return ColorYellow
	}

	if n >= CardGreenZero && n <= CardGreenDrawTwo {
		return ColorGreen
	}

	if n >= CardBlueZero && n <= CardBlueDrawTwo {
		return ColorBlue
	}

	if n >= CardWild && n <= CardWildDrawFour {
		return ColorSpecial
	}

	return NotColor
}
