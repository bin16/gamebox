package reversi

import "testing"

func TestCellsToReverse(t *testing.T) {
	g := NewGame()
	g.Status = GameStatusStarted

	c := Black
	n := "f4"
	x, y := g.indexOf(n)
	if x != 5 || y != 3 {
		t.Errorf("g.indexOf failed: %s -> %d,%d; want %d,%d", n, x, y, 5, 3)
	}

	cells := g.cellsToReserve(c, x, y)
	if len(cells) != 1 {
		t.Errorf("g.cellsToReverse failed: %d,%d;%d %v", x, y, c, cells)
	}
	g.Play(c, n)

	np := g.nextPlayer()
	if np != g.reverseOf(c) {
		t.Errorf("g.nextPlayer/g.reverseOf failed: nextPlayer is %d, want %d", np, g.reverseOf(c))
	}

	g.Play(np, "d3")
	rooms := g.getRooms(c)
	if len(rooms) != 5 {
		g.print()
		t.Errorf("g.Play may failed: %v", rooms)
	}
}
