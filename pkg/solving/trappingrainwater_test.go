package solving

import (
	"testing"
)

func TestTrappingRainWater (t *testing.T) {
	cases := []struct{
		name string
		heights []int
		expectedOutput int
		isErrorExpected bool
	} {
		{"Empty: single width elevation map", []int{0}, 0, false},
		{"Empty: 2 width elevation map, left lean", []int{1, 0}, 0, false},
		{"Empty: 2 width elevation map, right lean", []int{0, 1}, 0, false},
		{"Edge: left height, right no height", []int{1, 0, 0}, 0, false},
		{"Edge: right height, left no height", []int{0, 0, 1}, 0, false},
		{"Unit: left > right", []int{2, 0, 1}, 1, false},
		{"Unit: left < right", []int{1, 0, 2}, 1, false},
		{"Unit: left == right/valley", []int{1, 0, 1}, 1, false},
		{"Unit: hill", []int{1, 2, 1}, 0, false},
		{"Unit: right cliff", []int{1, 2, 0}, 0, false},
		{"Unit: left cliff", []int{0, 2, 1}, 0, false},
		{"Full: valley hill valley", []int{3, 0, 1, 2, 0, 0, 0, 4}, 3 + 2 + 1 + 3 + 3 + 3, false},
		{"Full: smooth valley", []int{3, 2, 1, 0, 0, 0, 1, 2, 3}, 1 + 2 + 3 + 3 + 3 + 2 + 1, false},
		{"Full: hill valley hill", []int{1, 2, 3, 2, 1, 0, 0, 0, 1, 2, 3, 2, 1}, 1 + 2 + 3 + 3 + 3 + 2 + 1, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := NewTrappingRainWaterProblem(c.heights)

			got, err := p.Solve()
			if err != nil {
				if !c.isErrorExpected {
					t.Errorf("Name: %v, Expected: no errors, Got: %+v", c.name, err)
				}
				return
			}

			if got.(trappingRainWaterSolution).TrappedWater != c.expectedOutput {
				t.Fatalf("Name: %v, Expected: %v, Got: %v", c.name, c.expectedOutput, got)
			}
		})
	}
}