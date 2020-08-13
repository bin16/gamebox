package reversi

import "testing"

func TestFull(t *testing.T) {
	a := "Alice"
	b := "Bob"
	g := NewGame()

	// Alice join
	if err := g.Join(a); err != nil {
		t.Errorf("Failed to run g.Join(%s), %s", a, err.Error())
	}
	if ca := g.SideOfPlayer(a); ca != Black {
		t.Errorf("Failed to run g.SideOfPlayer(%s), got %d, want %d", a, Black, ca)
	}
	if g.Status != GameStatusOpen {
		t.Errorf("Game status failed, got %s, want %s", g.Status, GameStatusOpen)
	}

	// Bob join
	if cb := g.SideOfPlayer(b); cb != Blank {
		t.Errorf("Failed to run g.SideOfPlayer(%s), got %d, want %d", b, Blank, cb)
	}
	if err := g.Join(b); err != nil {
		t.Errorf("Failed to run g.Join(%s), %s", b, err.Error())
	}
	if g.Status != GameStatusStarted {
		t.Errorf("Game status failed, got %s, want %s", g.Status, GameStatusStarted)
	}
	if r := g.playerInGame(b); !r {
		t.Errorf("Failed to run g.playerInGame(%s), got %v, want %v", b, r, true)
	}
	if cb := g.SideOfPlayer(b); cb != White {
		t.Errorf("Failed to run g.SideOfPlayer(%s), got %d, want %d", b, cb, White)
	}

	// TODO: more steps
}
