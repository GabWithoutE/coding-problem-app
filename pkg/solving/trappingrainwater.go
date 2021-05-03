package solving

// Problem: Given n non-negative ints representing elevation map
// 		compute how much water can be trapped.

type trappingRainWaterProblem struct {
	Heights []int
}

func NewTrappingRainWaterProblem(heights []int) Problem {
	return &trappingRainWaterProblem{
		Heights: heights,
	}
}

type trappingRainWaterSolution struct {
	TotalWater int
}

// TODO: Implement solving the Rainwater problem practicing DP
func (p *trappingRainWaterProblem) Solve() (interface{}, error) {
	return trappingRainWaterSolution{
		TotalWater: 0,
	}, nil
}