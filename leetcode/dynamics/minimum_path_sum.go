package dynamics

/**
64
https://leetcode.com/problems/minimum-path-sum/
 */

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return  0
	}
	rows := len(grid)
	cols := len(grid[0])
	//initial path
	path := make([][]int, rows)
	for i:=0; i<rows; i++ {
		path[i] = make([]int, cols)
	}

	path[0][0] = grid[0][0]
	for i:=1; i<rows; i++ {
		path[i][0] = path[i-1][0] + grid[i][0]
	}
	for j:=1; j<cols; j++ {
		path[0][j] = path[0][j-1] + grid[0][j]
	}

	for i:=1; i<rows; i++ {
		for j:=1; j<cols; j++ {
			path[i][j] = min(path[i-1][j], path[i][j-1])+grid[i][j]
		}
	}
	return path[rows-1][cols-1]
}

func min(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}
