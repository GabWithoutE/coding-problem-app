package solving

type Problem interface {
	Solve() (interface{}, error)
}

type problemsCatalogue struct{}

type ProblemsCatalogue interface {
	NewWordBreakProblem(unBrokenString string, wordDictionary []string) Problem
}

func NewProblemsCatalogue() ProblemsCatalogue {
	return &problemsCatalogue{}
}
