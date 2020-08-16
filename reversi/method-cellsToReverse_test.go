package reversi

import "testing"

func TestCellsToReverse(t *testing.T) {
	p1 := "Foo"
	p2 := "Bar"
	g := NewGame()
	g.Join(p1)
	g.Join(p2)
	g.Status = GameStatusStarted

	c := Black
	n := "f5"
	x, y := g.indexOf(n)
	assertIndexOf(t, g)("f5", 5, 4)

	cells := g.cellsToReserve(c, x, y)
	if len(cells) != 1 {
		g.print()
		t.Errorf("Failed to run g.cellsToReverse(%d, %d ,%d); got %v", c, x, y, cells)
	}

	mustPlay(t, g)(p1, n)
}
