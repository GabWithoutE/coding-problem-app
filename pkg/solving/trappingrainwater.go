package solving

// Problem: Given n non-negative ints representing elevation map
// 		compute how much water can be trapped.

type trappingRainWaterProblem struct {
	heights []int
}

func (p *problemsCatalogue) NewTrappingRainWaterProblem(heights []int) Problem {
	return &trappingRainWaterProblem{
		heights: heights,
	}
}

type trappingRainWaterSolution struct {
	TrappedWater int `json:"trappedWater"`
}

func (p *trappingRainWaterProblem) Solve() (interface{}, error) {
	// If the elevation map is only 2 wide, water trapping is impossible
	if len(p.heights) <= 2 {
		return trappingRainWaterSolution{
			TrappedWater: 0,
		}, nil
	}

	// Gives the max height at a given location relative to the left hand
	//		side for j = 0, and right hand for j = 1
	currentMaxHeightTable := make([][]int, 2)
	for i, _ := range currentMaxHeightTable {
		currentMaxHeightTable[i] = make([]int, len(p.heights))
	}

	currentMaxHeightTable[0][0] = p.heights[0]
	for i := 1; i < len(p.heights); i++ {
		h := p.heights[i]
		prevH := currentMaxHeightTable[0][i-1]
		if h < prevH {
			currentMaxHeightTable[0][i] = prevH
			continue
		}

		currentMaxHeightTable[0][i] = h
	}

	accumWaterTable := make([]int, len(p.heights))
	currentMaxHeightTable[1][len(p.heights)-1] = p.heights[len(p.heights)-1]
	for i := len(p.heights) - 2; i >= 0; i-- {
		h := p.heights[i]
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
