package dynamics

/**
120
https://leetcode.com/problems/triangle/
 */

func minimumTotal(triangle [][]int) int {
	for i:=1; i<len(triangle); i++ {
		for j:=0; j<len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
			}
		}
	}
	res := triangle[len(triangle)-1][0]
	for x:=1; x<len(triangle[len(triangle)-1]); x++ {
		res = min(res, triangle[len(triangle)-1][x])
	}
	return res
}

