package reversi

import "testing"

func TestIndexOf(t *testing.T) {
	g := NewGame()
	names := []string{
		"a1", "b2", "c8", "b6", "e5", "E6", "f3", "D6",
	}
	indexes := [][2]int{
		[2]int{0, 0},
		[2]int{1, 1},
		[2]int{2, 7},
		[2]int{1, 5},
		[2]int{4, 4},
		[2]int{4, 5},
		[2]int{5, 2},
		[2]int{3, 5},
	}

	for i, n := range names {
		r := indexes[i]
		x := r[0]
		y := r[1]
		assertIndexOf(t, g)(n, x, y)
	}
}
