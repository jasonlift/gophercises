package array

import "testing"

func TestPermutationSequence(t *testing.T) {
	res := getPermutation(3, 3)
	res2 := getPermutationII(3, 3)
	if res != "213" || res2 != "213" {
		t.Errorf("get %v", res)
	}

	res = getPermutation(3, 5)
	if res != "312" {
		t.Errorf("get %v", res)
	}
}
