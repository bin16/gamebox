package reversi

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	p1 := "Foo" // Black
	p2 := "Bar" // White
	g := NewGame()
	mustJoin(t, g)(p1)
	mustJoin(t, g)(p2)
	steps := `1 f5
2 d6
1 c3
2 d3
1 c4
2 f4
1 c5
2 b3
1 c2
2 e3
1 d2
2 c6
1 b4
2 b5
1 f2
2 e2
1 a3
2 e1
1 d1
2 f3
1 f1
2 e6
1 f6
2 g5
1 a6
2 b6
1 g3
2 g6
1 g4
2 h4
1 d7
2 e7
1 c7
2 b1
1 f8
2 b2
1 h3
2 h2
1 h6
2 f7
1 e8
2 h5
1 h1
2 c1
1 a7
2 c8
1 d8
2 b8
1 a8
2 b7
1 g2
2 g1
1 a4
2 a5
1 a1
2 a2
2 g8
1 h8
2 h7
1 g7`
	name := func(num string) string {
		if num == "1" {
			return p1
		}

		return p2
	}
	for _, s := range strings.Split(steps, "\n") {
		r := strings.Split(s, " ")
		mustPlay(t, g)(name(r[0]), r[1])
	}
	assertScoreIs(t, g)([3]int{39, 25, 64})
}
