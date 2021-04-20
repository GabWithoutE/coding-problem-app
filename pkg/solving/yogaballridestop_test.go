package solving

import (
	"testing"
)

func TestYogaballRideStopSolver (t *testing.T) {
	cases := []struct{
		name string
		startPos int
		startSpeed int
		runway []bool
		expectedOutput bool
		isErrorExpected bool
	} {
		{"start on a pin", 0, 1, []bool{false, true, true}, false, false},
		{"start on non 0 index", 1, 1, []bool{false, true, true}, true, false},
		{"start on last index", 2, 1, []bool{false, true, true}, true, false},
		{"runway length 1", 0, 1, []bool{true}, true, false},
		{"over speed: > runway length", 0, 4, []bool{true, true, true}, false, false},
		{"over speed: sum(0, initS) > runway", 0, 3, []bool{true, true, false, true, true}, false, false},
		{"normal: cannot stop", 0, 2, []bool{true, true, false, true, false}, false, false},
		{"normal: can stop", 0, 2, []bool{true, true, false, true, true}, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := NewYogaballRideStopProblem(c.startPos, c.startSpeed, c.runway)

			got, err := p.Solve()
			if err != nil {
				if !c.isErrorExpected {
					t.Errorf("Name: %v, Expected: no errors, Got: %+v", c.name, err)
				}
				return
			}

			if got.(yogaballRideStopSolution).IsStoppable != c.expectedOutput {
				t.Fatalf("Name: %v, Expected: %v, Got: %v", c.name, c.expectedOutput, got)
			}
		})
	}
}
