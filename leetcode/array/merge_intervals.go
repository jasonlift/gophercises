package array

import "sort"

/**
56
https://leetcode.com/problems/merge-intervals/
 */

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	if len(intervals) == 0 {
		return res
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	cur := intervals[0]
	for i:=1; i<len(intervals); i++ {
		tmp := intervals[i]
		if tmp[0] >= cur[0] && tmp[0] <= cur[1] {
			end := max(cur[1], tmp[1])
			cur = []int{cur[0], end}
			continue
		}
		res = append(res, cur)
		cur = tmp
	}
	res = append(res, cur)
	return res
}

func max(x, y int) int {if x > y {return x}; return y}