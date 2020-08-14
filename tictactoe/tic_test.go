package tictactoe

import (
	"fmt"
	"testing"
)

func TestTicTacToe(t *testing.T) {
	p1 := "Foo"
	p2 := "Bar"
	g := NewGame()
	mustJoin(t, g)(p1)
	mustJoin(t, g)(p2)
	mustPlay(t, g)(p1, "b2")
	mustPlay(t, g)(p2, "c3")
}

func TestTools(t *testing.T) {
	g := NewGame()
	if x, y := g.indexOf("c2"); x != 2 || y != 1 {
		t.Errorf("Failed to run g.indexOf(%s); got %d,%d; want %d,%d", "c3", x, y, 2, 1)
	}

	if n := g.nameOf(0, 2); n != "a3" {
		t.Errorf("Failed to run g.nameOf(%d, %d); got %s; want %s", 0, 2, n, "a3")
	}

	if ns := g.nextSide(); ns != Blank {
		t.Errorf("Failed to run g.nextSide(); got %d; want %d; %v", ns, Blank, g.History)
	}
	g.Status = StatusStarted
	g.History = [][3]int{
		[3]int{int(X), 0, 0},
	}
	if ns := g.nextSide(); ns != O {
		t.Errorf("Failed to run g.nextSide(); got %d; want %d; %v", ns, O, g.History)
	}
	g.History = [][3]int{
		[3]int{int(O), 0, 0},
	}
	if ns := g.nextSide(); ns != X {
		t.Errorf("Failed to run g.nextSide(); got %d; want %d; %v", ns, X, g.History)
	}
	g.History = [][3]int{
		[3]int{int(X), 0, 0},
		[3]int{int(O), 0, 0},
	}
	if ns := g.nextSide(); ns != X {
		t.Errorf("Failed to run g.nextSide(); got %d; want %d; %v", ns, X, g.History)
	}

}

func (g *Game) print() {
	text := ""
	for y := 0; y < boardSize; y++ {
		row := ""
		for x := 0; x < boardSize; x++ {
			c := g.Board[x][y]
			row += piece(c)
		}
		text += row + "\n"
	}

	logText := ""
	for i, log := range g.History {
		x := log[1]
		y := log[2]
		logText += fmt.Sprintf("0%d. %s - %s", i+1, piece(Side(log[0])), g.nameOf(x, y))
	}
	fmt.Printf(`====
Status: %s
Players: %v
Next: %s (%d)
----
%s----
%s
====
`, g.statusDict()[g.Status], g.Players, piece(g.nextSide()), g.nextSide(), text, logText)
}

func piece(c Side) string {
	l := [3]string{"⬜", "❌", "⭕️"}
	return l[int(c)]
}

func mustPlay(t *testing.T, g *Game) func(string, string) {
	return func(playerID, cmd string) {
		if err := g.Play(playerID, cmd); err != nil {
			t.Errorf("Failed to run g.Play(%s, %s), got error <%v>", playerID, cmd, err)
			g.print()
		}
	}
}

func mustJoin(t *testing.T, g *Game) func(string) {
	return func(playerID string) {
		if err := g.Join(playerID); err != nil {
			t.Errorf("Failed to run g.Join(%s), got error <%v>", playerID, err)
			g.print()
		}
	}
}
