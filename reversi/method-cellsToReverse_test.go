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
	if x != 5 || y != 4 {
		t.Errorf("Failed to run g.indexOf(%s); got %d,%d; want %d,%d", n, x, y, 5, 3)
	}

	cells := g.cellsToReserve(c, x, y)
	if len(cells) != 1 {
		g.print()
		t.Errorf("Failed to run g.cellsToReverse(%d, %d ,%d); got %v", c, x, y, cells)
	}

	if err := g.Play(p1, n); err != nil {
		g.print()
		t.Errorf("Failed to run g.Play(%s, %s); got error %v", p1, n, err)
	}
}
