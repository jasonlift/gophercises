package strings

import "strings"

/**
151
https://leetcode.com/problems/reverse-words-in-a-string/
 */

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	tmp := []string{}
	for i := range words {
		if words[i] == "" {
			continue
		}
		covWord := swap(words[i])
		tmp = append(tmp, covWord)
	}
	return swap(strings.Join(tmp, " "))
}

func swap(s string) string {
	sli := []byte(s)
	left, right := 0, len(sli)-1
	for left < right {
		sli[left], sli[right] = sli[right], sli[left]
		left++
		right--
	}
	return string(sli)
}