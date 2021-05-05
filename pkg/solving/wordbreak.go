package solving

type wordBreakProblem struct {
	UnbrokenString string
	WordDictionary []string
}

func NewWordBreakProblem(unbrokenString string, wordDictionary []string) Problem {
	return &wordBreakProblem{
		UnbrokenString: unbrokenString,
		WordDictionary: wordDictionary,
	}
}

type wordBreakSolution struct {
	SpacedStrings []string `json:"brokenStrings"`
}

func (p *wordBreakProblem) Solve() (interface{}, error) {

	return &wordBreakSolution{

	}, nil
}