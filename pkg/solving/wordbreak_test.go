package solving

import "testing"

// Problem: Given a string, and a dictionary of words, add spaces to s to create a valid set of dictionary words.
//		Return all valid sets in any order as strings

type orderlessComparissonStringSet map[string]bool

func (oss *orderlessComparissonStringSet) Equals(comp []string) bool {
	if len(*oss) != len(comp) {
		return false
	}
//	TODO: Finish implementing check for all permutations
}

func TestWordBreak(t *testing.T) {
	_ := []struct {
		name            string
		unbrokenString  string
		wordDictionary  []string
		expectedOutput  orderlessComparissonStringSet
		isErrorExpected bool
	}{
		{},
	}
// TODO: Finish writing tests...

}
