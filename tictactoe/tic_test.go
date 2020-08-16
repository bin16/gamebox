package tictactoe

import (
	"fmt"
	"testing"
)

func TestWinCheck(t *testing.T) {
	p1 := "Foo"
	p2 := "Bar"
	g := NewGame()
	mustJoin(t, g)(p1)
	mustJoin(t, g)(p2)

	winnerIs(t, g)(Blank)

	g.Board = [boardSize][boardSize]Side{
		[boardSize]Side{X, O, Blank},
		[boardSize]Side{X, O, Blank},
		[boardSize]Side{X, Blank, Blank},
	}
	winnerIs(t, g)(X)
	lineIsSides(t, g)(0, 0, 1, 0, [boardSize]Side{X, X, X})
	lineIsSides(t, g)(0, 0, 1, 1, [boardSize]Side{X, O, Blank})
	lineIsSides(t, g)(0, 0, 0, 1, [boardSize]Side{X, O, Blank})
	pieceIsWin(t, g)(0, 0, true)

	lineIsSides(t, g)(1, 0, 0, 1, [boardSize]Side{X, O, Blank})
	pieceIsWin(t, g)(1, 0, true)

	lineIsSides(t, g)(2, 0, 0, 1, [boardSize]Side{X, Blank, Blank})
	pieceIsWin(t, g)(2, 0, true)

	pieceIsWin(t, g)(0, 1, false)
	pieceIsWin(t, g)(1, 1, false)
	pieceIsWin(t, g)(2, 1, false)

	g.Board = [boardSize][boardSize]Side{
		[boardSize]Side{X, O, Blank},
		[boardSize]Side{O, X, Blank},
		[boardSize]Side{X, Blank, X},
	}
	winnerIs(t, g)(X)

	g.Board = [boardSize][boardSize]Side{
		[boardSize]Side{X, O, O},
		[boardSize]Side{O, O, Blank},
		[boardSize]Side{O, Blank, X},
	}
	winnerIs(t, g)(O)
}

func TestTicTacToe(t *testing.T) {
	p1 := "Foo"
	p2 := "Bar"
	g := NewGame()
	mustJoin(t, g)(p1)
	mustJoin(t, g)(p2)
	gameIs(t, g)(StatusStarted)
	mustPlay(t, g)(p1, "b2") // 1
	gameIs(t, g)(StatusStarted)
	winnerIs(t, g)(Blank)
	mustPlay(t, g)(p2, "c3") // 2
	mustPlay(t, g)(p1, "c2") // 3
	mustPlay(t, g)(p2, "a2") // 4
	mustPlay(t, g)(p1, "a3") // 5
	mustPlay(t, g)(p2, "b1") // 6
	mustPlay(t, g)(p1, "a1") // 7
	winnerIs(t, g)(Blank)
	mustPlay(t, g)(p2, "c1") // 8
	gameIsEnd(t, g)(false, Blank)
	winnerIs(t, g)(Blank)
	mustPlay(t, g)(p1, "b3") // 9
	gameIsEnd(t, g)(true, Blank)
	gameIs(t, g)(StatusEnded)
	winnerIs(t, g)(Blank)

	g2 := NewGame()
	gameIs(t, g2)(StatusOpen)
	mustJoin(t, g2)(p1)
	gameIs(t, g2)(StatusOpen)
	mustJoin(t, g2)(p2)
	gameIs(t, g2)(StatusStarted)
	mustPlay(t, g2)(p1, "b2") // 1
	gameIs(t, g2)(StatusStarted)
	winnerIs(t, g2)(Blank)
	mustPlay(t, g2)(p2, "c3") // 2
	winnerIs(t, g2)(Blank)
	mustPlay(t, g2)(p1, "c2") // 3
	winnerIs(t, g2)(Blank)
	mustPlay(t, g2)(p2, "b3") // 4
	winnerIs(t, g2)(Blank)
	mustPlay(t, g2)(p1, "a2") // 5
	gameIsEnd(t, g2)(true, X)
	gameIs(t, g2)(StatusEnded)
	winnerIs(t, g2)(X)
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
		logText += fmt.Sprintf("0%d. %s - %s\n", i+1, piece(Side(log[0])), g.nameOf(x, y))
	}
	fmt.Printf(`====
Status: %s
Players: %v
Next: %s (Side=<%d>)
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
		t.Helper()
		if err := g.Play(playerID, cmd); err != nil {
			t.Errorf("Failed to run g.Play(%s, %s), got error <%v>", playerID, cmd, err)
			g.print()
		}
	}
}

func mustJoin(t *testing.T, g *Game) func(string) {
	return func(playerID string) {
		t.Helper()
		if err := g.Join(playerID); err != nil {
			t.Errorf("Failed to run g.Join(%s), got error <%v>", playerID, err)
			g.print()
		}
	}
}

func gameIs(t *testing.T, g *Game) func(Status) {
	return func(s Status) {
		t.Helper()
		if g.Status != s {
			t.Errorf("Failed to check g.Status, got %d, want %d", g.Status, s)
			g.print()
		}
	}
}

func winnerIs(t *testing.T, g *Game) func(Side) {
	return func(c Side) {
		t.Helper()
		if g.winnerSide() != c {
			t.Errorf("Failed to run g.winnerSide(), got %d, want %d", g.winnerSide(), c)
			g.print()
		}
	}
}

func lineIsSides(t *testing.T, g *Game) func(int, int, int, int, [boardSize]Side) {
	return func(x0, y0, dx, dy int, l [boardSize]Side) {
		t.Helper()
		l1 := g.lineOf(x0, y0, dx, dy, boardSize)
		for i, p := range l1 {
			if p != l[i] {
				t.Errorf("Failed to run g.lineOf(%d, %d, %d, %d, %d); want %v; got %v", x0, y0, dx, dy, boardSize, l, l1)
			}
		}
	}
}

func pieceIsWin(t *testing.T, g *Game) func(int, int, bool) {
	return func(x, y int, b bool) {
		t.Helper()
		if g.isWin(x, y) != b {
			h := g.lineOf(0, y, 1, 0, boardSize)
			v := g.lineOf(x, 0, 0, 1, boardSize)
			t.Errorf("Failed to run g.isWin(%d, %d), got %v, want %v; \nH: %v;\nV: %v;\nBoard:%v", x, y, g.isWin(x, y), b, h, v, g.Board)
			g.print()
		}
	}
}

func gameIsEnd(t *testing.T, g *Game) func(bool, Side) {
	return func(e bool, c Side) {
		t.Helper()
		if e1, c1 := g.End(); e1 != e || c1 != c {
			t.Errorf("Failed to run g.End(); got %v,%v; want %v,%v", e1, c1, e, c)
		}
	}
}
