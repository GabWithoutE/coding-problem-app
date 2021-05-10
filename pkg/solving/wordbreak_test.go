package solving

import (
	"testing"
)

type stringSet map[string]struct{}


func (oss stringSet) Equals(comp []string) bool {
	if len(oss) != len(comp) {
		return false
	}

	if len(oss) == 0 && len(comp) == 0 {
		return true
	}

	// Check to make sure all of the values in comp are present in the stringSet
	for _, s := range comp {
		_, exist := oss[s]
		if !exist {
			return false
		}
	}

	return true
}

var (
	exists = struct{}{}
	emptySet = stringSet{}
)

func TestWordBreak(t *testing.T) {
	cases := []struct {
		name            string
		unbrokenString  string
		wordDictionary  []string
		expectedOutput  stringSet
		isErrorExpected bool
	}{
		{"Empty: word dictionary empty", "c", make([]string, 0), emptySet, false},
		{"Empty: empty input string", "", []string{"cap"}, emptySet, false},
		{"Unit: match 1 letter, 1 dict", "c", []string{"c"}, stringSet{"c": exists}, false},
		{"Unit: match 2 ways", "ca", []string{"ca", "c", "a"}, stringSet{"c a": exists, "ca": exists}, false},
		{"Unit: fail partial match", "cab", []string{"ca", "c", "a"}, emptySet, false},
		{"Full: match sub, half, and full", "cabcap", []string{"cab", "cap", "ca", "b", "p", "cabcap"}, stringSet{"cabcap": exists, "cab ca p": exists, "cab cap": exists, "ca b ca p": exists, "ca b cap": exists}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ps := NewProblemsCatalogue()
			p := ps.NewWordBreakProblem(c.unbrokenString, c.wordDictionary)

			got, err := p.Solve()
			if err != nil {
				if !c.isErrorExpected {
					t.Errorf("Name: %v, Expected: no errors, Got: %+v", c.name, err)
				}
				return
			}

			if !c.expectedOutput.Equals(got.(wordBreakSolution).SpacedStrings) {
				t.Fatalf("Name: %v, Expected: %v, Got: %v", c.name, c.expectedOutput, got)
			}
		})
	}
}
