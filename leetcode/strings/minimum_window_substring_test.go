package strings

import "testing"

func TestMinimumWindow(t *testing.T) {
	S := "ADOBECODEBANC"
	T := "ABC"
	desired := "BANC"

	res := minWindow(S, T)
	if res != desired {
		t.Errorf("wrong answer: %v", res)
	}
}
