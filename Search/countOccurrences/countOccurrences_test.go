package countOccurrences

import "testing"

func TestCount(t *testing.T) {

	cases := []struct {
		in           []int
		search, want int
	}{
		{[]int{0, 1, 1, 3, 3, 3, 3, 6, 8, 10, 11, 11}, 3, 4},
		{[]int{0, 1, 1, 3, 6, 8, 10, 11, 11}, 3, 1},
		{[]int{0, 1, 1, 6, 8, 10, 11, 11}, 3, 0},
		{[]int{3, 3, 3, 3, 3}, 3, 5},
		{[]int{3}, 3, 1},
		{[]int{0}, 3, 0},
	}

	for _, c := range cases {
		got := Count(c.in, c.search)

		if got != c.want {
			t.Errorf("Count(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
