package strings

import "strings"

/**
71
https://leetcode.com/problems/simplify-path/
 */

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	stack := make([]string, 0, len(path))
	// 0 is specified length, len(path) is capacity of slice
	for _, d := range dirs {
		if len(d) == 0 || d == "." {
			continue
		}

		if d == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if d != ".." {
			stack = append(stack, d)
		}
	}
	return "/" + strings.Join(stack, "/")
}
