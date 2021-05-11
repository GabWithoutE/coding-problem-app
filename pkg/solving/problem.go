package solving

type Problem interface {
	Solve() (interface{}, error)
}

type problemsCatalogue struct{}

type ProblemsCatalogue interface {
	NewWordBreakProblem(unBrokenString string, wordDictionary []string) Problem
	NewCoinChangeProblem(denominations []int, total int) Problem
	NewTrappingRainWaterProblem(heights []int) Problem
	NewYogaballRideStopProblem(startPos int, startSpeed int, runway []bool) Problem
}

func NewProblemsCatalogue() ProblemsCatalogue {
	return &problemsCatalogue{}
}
