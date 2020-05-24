package strings

import (
	"strconv"
	"strings"
)

/**
93
https://leetcode.com/problems/restore-ip-addresses/
 */

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	dfs(&res, s, make([]string, 0), 0)
	return res
}

func dfs(res *[]string, s string, str []string, index int) {
	if len(str) == 4 {
		if index == len(s) { // complete the process of dfs
			*res = append(*res, strings.Join(str, "."))
		}
		return
	}
	// index means which position of char in the s
	for i := index; i < index + 3 && i < len(s); i++ {
		candidate := s[index: i+1]
		if num,_ := strconv.Atoi(candidate); num > 255 {
			break
		}

		str = append(str, candidate)
		dfs(res, s, str, i + 1)
		str = str[:len(str) - 1]

		// why break?
		// because 0 in this segment is the only one choice
		// 192.168.0.18, no 192.168.01.8
		if candidate == "0" {
			break
		}
	}
}