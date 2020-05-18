package strings

/**
https://leetcode.com/problems/longest-common-prefix/

 */

/**
use partition and merge

 */
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	return partition(strs, 0, len(strs)-1)
}

func partition(strs []string, s int, e int) string {
	if s == e {
		return strs[s]
	} else if e - s > 1 {
		mid := (e+s) / 2
		s1 := partition(strs, s, mid)
		s2 := partition(strs, mid+1, e)
		return commonPrefix(s1, s2)
	} else {
		return commonPrefix(strs[s], strs[e])
	}
}

func commonPrefix(s1 string, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}

	i := 0
	for i < len(s1) {
		if len(s2) <= i {
			return s1[0: i]
		}

		if s1[i] != s2[i] {
			return s1[0: i]
		} else {
			i++
		}

		if i == len(s1) {
			return s1[0: i]
		}
	}
	return ""
}
