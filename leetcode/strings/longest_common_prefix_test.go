package strings

import (
	"testing"
)

func TestGeneral(t *testing.T) {
	strs := []string{"flower","flow","flight"}

	res := longestCommonPrefix(strs)

	if "fl" != res {
		t.Error("wrong result, hopes \"fl\"")
	}
}

func TestTwoChars(t *testing.T) {
	strs := []string{"c","c"}

	res := longestCommonPrefix(strs)

	if "c" != res {
		t.Error("wrong result, hopes \"c\"")
	}
}