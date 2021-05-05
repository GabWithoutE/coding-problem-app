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
	TrappedWater int `json:"TrappedWater"`
}

func (p *trappingRainWaterProblem) Solve() (interface{}, error) {
	// If the elevation map is only 2 wide, water trapping is impossible
	if len(p.Heights) <= 2 {
		return trappingRainWaterSolution{
			TrappedWater: 0,
		}, nil
	}

	// Gives the max height at a given location relative to the left hand
	//		side for j = 0, and right hand for j = 1
	currentMaxHeightTable := make([][]int, 2)
	for i, _ := range currentMaxHeightTable {
		currentMaxHeightTable[i] = make([]int, len(p.Heights))
	}

	currentMaxHeightTable[0][0] = p.Heights[0]
	for i := 1; i < len(p.Heights); i++ {
		h := p.Heights[i]
		prevH := currentMaxHeightTable[0][i-1]
		if h < prevH {
			currentMaxHeightTable[0][i] = prevH
			continue
		}

		currentMaxHeightTable[0][i] = h
	}

	accumWaterTable := make([]int, len(p.Heights))
	currentMaxHeightTable[1][len(p.Heights)-1] = p.Heights[len(p.Heights)-1]
	for i := len(p.Heights) - 2; i >= 0; i-- {
		h := p.Heights[i]
		prevH := currentMaxHeightTable[1][i+1]

		if prevH > h {
			currentMaxHeightTable[1][i] = prevH
		} else {
			currentMaxHeightTable[1][i] = h
		}

		if currentMaxHeightTable[1][i] > currentMaxHeightTable[0][i] {
			accumWaterTable[i] = currentMaxHeightTable[0][i] - h + accumWaterTable[i+1]
			continue
		}

		accumWaterTable[i] = currentMaxHeightTable[1][i] - h + accumWaterTable[i+1]
	}

	return trappingRainWaterSolution{
		TrappedWater: accumWaterTable[0],
	}, nil
}
