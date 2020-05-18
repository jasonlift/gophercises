package strings

/**
https://leetcode.com/problems/longest-substring-without-repeating-characters/

 */

// not very correct
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	dic := make(map[string]int)
	longestLength := 1
	left := 0
	for i, c := range s {
		index, ok := dic[string(c)]
		if !ok {
			dic[string(c)] = i
			if i-left+1 > longestLength {
				longestLength = i-left+1
			}
		} else {
			dic[string(c)] = i
			if index < left {
				// contains itself
				if i-left+1 > longestLength {
					longestLength = i-left+1
				}
			} else {
				// eliminate itself
				if i-left > longestLength {
					longestLength = i - left
				}
			}
			left = index+1
		}
	}
	return longestLength
}