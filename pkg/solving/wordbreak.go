package solving

import "github.com/pkg/errors"

type wordBreakProblem struct {
	UnbrokenString string
	WordDictionary []string
}

// Problem: Given a string, and a dictionary of words, add spaces to s to create a valid set of dictionary words.
//		Return all valid sets in any order as strings

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

	return p.dp()
}

func (p *wordBreakProblem) dp() (interface{}, error) {
	if p.UnbrokenString == "" || len(p.WordDictionary) == 0 {
		return wordBreakSolution{
			SpacedStrings: make([]string, 0),
		}, nil
	}

	dict := make(map[string]bool)

	// initialize the word dictionary as a hash "set" (map in golang)
	for _, w := range p.WordDictionary {
		if _, exists := dict[w]; exists {
			return nil, &SolverError{
				Problem: "Word Break",
				Err: errors.New("Solve: duplicate dictionary entries"),
			}
		}
		dict[w] = true
	}

	breakpointTable := make([][]string, len(p.UnbrokenString))

	// initialize string arrays
	for bp, _ := range breakpointTable {
		breakpointTable[bp] = make([]string, 0)
	}

	// initialize first entry, at breakpoint 0
	s := p.UnbrokenString[0 : 1]
	if _, exists := dict[s]; exists {
		breakpointTable[0] = append(breakpointTable[0], s)
	}

	for bp := 1; bp < len(breakpointTable); bp++ {
		currentSubstring := p.UnbrokenString[0 : bp + 1]

		// check if the entire current substring is matchable
		if _, exists := dict[currentSubstring]; exists {
			breakpointTable[bp] = append(breakpointTable[bp], currentSubstring)
		}

		// check if can build upon previous substrings found @ prev breakpoints
		for i := 0; i < bp; i++ {
			// if no previous matches at bp, nothing to build on...
			if len(breakpointTable[i]) == 0 {
				continue
			}

			// remove the matched prefix, to check if postfix matches
			postfix := currentSubstring[i + 1 : bp + 1]

			// no match, then move on
			if _, exists := dict[postfix]; !exists {
				continue
			}

			// append the matches to the values
			for _, val := range breakpointTable[i] {
				breakpointTable[bp] = append(breakpointTable[bp], val + " " + postfix)
			}
		}
	}

	return wordBreakSolution{
		SpacedStrings: breakpointTable[len(p.UnbrokenString) - 1],
	}, nil
}