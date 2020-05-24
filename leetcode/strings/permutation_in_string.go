package strings

/**
567
https://leetcode.com/problems/permutation-in-string/
 */

// https://leetcode.com/problems/permutation-in-string/discuss/102623/Golang-solution
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 || len(s1) > len(s2) {
		return false
	}

	cnt1, cnt2, diff, ans := [26]int{}, [26]int{}, [26]int{}, [26]int{}
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	for i:=0; i < len(diff); i++ {
		diff[i] = cnt2[i]-cnt1[i]
	}
	s1len := len(s1)
	for j:=s1len; j < len(s2); j++ {
		if diff == ans {
			return true
		}
		diff[s2[j-s1len]-'a']--
		diff[s2[j]-'a']++
	}
	return diff == ans
}

// https://leetcode.com/problems/permutation-in-string/discuss/159258/Golang-solution
func checkInclusionV2(s1 string, s2 string) bool {
	cnt1, cnt2 := [26]int{}, [26]int{}
	for _, c := range s1 {
		cnt1[int(c-'a')]++
	}

	start := 0
	for i, c := range s2 {
		cnt2[int(c-'a')]++
		for start <= i && cnt2[int(s2[start]-'a')] > cnt1[int(s2[start]-'a')] {
			cnt2[int(s2[start]-'a')]--
			start++
		}
		if cnt1 == cnt2 {
			return true
		}
	}
	return cnt1 == cnt2
}