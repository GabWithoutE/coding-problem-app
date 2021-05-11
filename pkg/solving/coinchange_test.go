package solving

import (
	"testing"
)

func TestCoinChangeSolver(t *testing.T) {
	cases := []struct {
		name            string
		denominations   []int
		total           int
		expectedOutput  int
		isErrorExpected bool
	}{
		{"normal", []int{1, 2, 5}, 11, 3, false},
		{"total 0", []int{1, 2, 5}, 0, 0, false},
		{"denoms greater than total", []int{10}, 1, -1, false},
		{"denoms don't fit", []int{5}, 11, -1, false},
		{"no denoms", nil, 3, -1, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cat := NewProblemsCatalogue()
			p := cat.NewCoinChangeProblem(c.denominations, c.total)

			got, err := p.Solve()
			if err != nil {
				if !c.isErrorExpected {
					t.Errorf("Name: %v, Expected: no errors, Got: %+v", c.name, err)
				}
				return
			}

			if got.(coinChangeSolution).Coins != c.expectedOutput {
				t.Fatalf("Name: %v, Expected: %v, Got: %v", c.name, c.expectedOutput, got)
			}
		})
	}
}
