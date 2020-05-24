package strings

/**
https://leetcode.com/problems/longest-substring-without-repeating-characters/

 */

// not very correct
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	dict := make(map[rune]int)
	longestLength := 0
	left := 0
	for i, c := range s {
		index, ok := dict[c]
		if ok && index >= left {
			if i - left > longestLength {
				longestLength = i - left
			}
			left = index + 1
		}
		dict[c] = i
	}
	if len(s) - left > longestLength {
		longestLength = len(s) - left
	}
	return longestLength
}