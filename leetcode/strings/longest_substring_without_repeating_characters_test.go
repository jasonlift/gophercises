package strings

import "testing"

func TestLongestSubstring1(t *testing.T) {
	input := "abcabcbb"
	res := lengthOfLongestSubstring(input)
	if res != 3 {
		t.Errorf("want 3, get %d", res)
	}
}

func TestBlank(t *testing.T) {
	input := " "
	res := lengthOfLongestSubstring(input)
	if res != 1 {
		t.Errorf("want 1, get %d", res)
	}
}

func TestLongestSubstring2(t *testing.T) {
	input := "abba"
	res := lengthOfLongestSubstring(input)
	if res != 2 {
		t.Errorf("want 2, get %d", res)
	}
}

func TestLongestSubstring3(t *testing.T) {
	input := "tmmzuxt"
	res := lengthOfLongestSubstring(input)
	if res != 5 {
		t.Errorf("want 5, get %d", res)
	}
}

