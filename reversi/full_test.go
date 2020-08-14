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
	if g.Status != GameStatusOpen {
		t.Errorf("Game status failed, got %d, want %d", g.Status, GameStatusOpen)
	}

	// Bob join
	if err := g.Join(b); err != nil {
		t.Errorf("Failed to run g.Join(%s), %s", b, err.Error())
	}
	if g.Status != GameStatusStarted {
		t.Errorf("Game status failed, got %d, want %d", g.Status, GameStatusStarted)
	}

	// TODO: more steps
}
