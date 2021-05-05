package solving

import (
	"github.com/gabriellukechen/coding-problem-app/pkg/helpers"
	"github.com/pkg/errors"
)

// Problem: Given coins of integer denominations, and a total amount,
// 		compute the fewest number of coins whose values when summed is
// 		exactly equal to total.

type coinChangeProblem struct {
	Denominations []int
	Total         int
}

func NewCoinChangeProblem(denominations []int, total int) Problem {
	return &coinChangeProblem{
		Denominations: denominations,
		Total:         total,
	}
}

type coinChangeSolution struct {
	Coins int `json:"coins"`
}

func (p *coinChangeProblem) Solve() (interface{}, error) {
	if p.Denominations == nil {
		return nil, &SolverError{
			Problem: "Coin Change",
			Err:     errors.New("Solve: no denominations given"),
		}
	}

	cs := solveCoinChange(p.Denominations, p.Total)

	return coinChangeSolution{
		Coins: cs,
	}, nil
}

func solveCoinChange(denominations []int, total int) int {
	dp := make([]int, total+1)
	dp[0] = 0

	for i := 1; i <= total; i++ {
		var counts []int

		for _, cval := range denominations {
			if i-cval >= 0 && dp[i-cval] >= 0 {
				counts = append(counts, dp[i-cval])
			}
		}

		if len(counts) < 1 {
			dp[i] = -1
			continue
		}

		dp[i] = helpers.Minimum(counts) + 1
	}

	return dp[total]
}
