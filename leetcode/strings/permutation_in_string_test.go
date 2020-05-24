package strings

import "testing"

func TestPermutationString(t *testing.T) {
	s1 := "ky"
	s2 := "ainwkckifykxlribaypk"

	if !checkInclusion(s1, s2) {
		t.Error("wrong answer")
	}
}
