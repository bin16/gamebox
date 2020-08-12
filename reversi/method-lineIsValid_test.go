package reversi

import "testing"

func TestLineIsValid(t *testing.T) {
	g := NewGame()
	l1 := []Side{
		Blank, White, White, Black, Blank,
	}
	if c := Black; g.lineIsValid(Black, l1) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l1)
	}
	if c := White; g.lineIsValid(White, l1) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l1)
	}

	l2 := []Side{
		Black, Black, Blank, White,
	}
	if c := Black; g.lineIsValid(Black, l2) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l2)
	}
	if c := White; g.lineIsValid(White, l2) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l2)
	}

	l3 := []Side{
		Black, Black, White,
	}
	if c := Black; g.lineIsValid(Black, l3) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l3)
	}
	if c := White; !g.lineIsValid(White, l3) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l3)
	}

	l4 := []Side{
		Black, Black,
	}
	if c := Black; g.lineIsValid(Black, l4) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l4)
	}
	if c := White; g.lineIsValid(White, l4) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l4)
	}

	l5 := []Side{
		Black,
	}
	if c := Black; g.lineIsValid(Black, l5) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l5)
	}
	if c := White; g.lineIsValid(White, l5) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l5)
	}

	l6 := []Side{}
	if c := Black; g.lineIsValid(Black, l6) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l6)
	}
	if c := White; g.lineIsValid(White, l6) {
		t.Errorf("g.lineIsValid failed: %d, %v", c, l6)
	}
}
