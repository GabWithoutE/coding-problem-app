package solving

// Recurrence Relation
// 		P is position of the ball, S is speed of the ball
// 		CS(P,S) = CS(P+S, S) || CS(P+S+1, S+1) || CS(P+S-1, S-1)
// 		Target problem is P(Start Position, Start Speed), so the subproblems should be done backwards
//
// Table Formulation
// 		Table Size: x, y = runway size - start P
//		Y speed dimension should be set because if speed is higher than runway - start Pos,
// 			it is guaranteed that the ball cannot stop
//			technically actually even x + x-1 ... 0 <= runway size - start P where x is start speed
//				if x does not fullfill the above inequality, then x is out of bounds...
//				x * (x + 1) / 2 <= runway size - start P
//
//		Table Example:
//      CS(P=0, S=2)
//		Can flip P axis to make implementation easier (index will match position)
//		  P 4 3 2 1 0
//      S -----------
//		0 | Y Y N Y Y
//		1 | Y Y N Y Y
//		2 | N Y N Y Y <- target problem
//		3 | N N N Y
// 		4 | N N N N

type yogaballRideStopProblem struct {
	startPos   int
	startSpeed int
	runway     []bool
}

func (p *problemsCatalogue) NewYogaballRideStopProblem(startPos int, startSpeed int, runway []bool) Problem {
	return &yogaballRideStopProblem{
		startPos:   startPos,
		startSpeed: startSpeed,
		runway:     runway,
	}
}

type yogaballRideStopSolution struct {
	IsStoppable bool `json:"isStoppable"`
}

func (p *yogaballRideStopProblem) Solve() (interface{}, error) {
	sol := yogaballRideStopSolution{}

	err := p.DP(&sol)

	return sol, err
}

func (p *yogaballRideStopProblem) DP(sol *yogaballRideStopSolution) error {
	runway := p.runway[p.startPos:]
	end := len(runway) - 1
	ss := p.startSpeed

	if p.startSpeed == 1 {
		ss = 0
	}

	if float32(ss*(ss+1))/float32(2) > float32(end) {
		sol.IsStoppable = false
		return nil
	}

	// Create DP table
	subProbsTable := make([][]bool, len(runway))
	for i, _ := range subProbsTable {
		subProbsTable[i] = make([]bool, len(runway))
	}

	for pos := end; pos >= 0; pos-- {
		for speed := 0; speed < len(runway); speed++ {
			if !runway[pos] {
				subProbsTable[pos][speed] = false
				continue
			}
			if speed == 0 || speed == 1 {
				subProbsTable[pos][speed] = true
				continue
			}
			if pos+speed > end || pos+speed+1 > end || pos+speed-1 > end {
				subProbsTable[pos][speed] = false
				continue
			}

			subProbsTable[pos][speed] = subProbsTable[pos+speed][speed] || subProbsTable[pos+speed+1][speed+1] || subProbsTable[pos+speed-1][speed-1]
		}
	}

	sol.IsStoppable = subProbsTable[0][ss]

	return nil
}
