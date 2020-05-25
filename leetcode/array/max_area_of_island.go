package array

/**
695
https://leetcode.com/problems/max-area-of-island/

solution：
https://www.cnblogs.com/coding-gaga/p/11241283.html
 */

var directions = [4][2]int{
	[2]int{0, 1},
	[2]int{0, -1},
	[2]int{1, 0},
	[2]int{-1, 0},
}

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	max := 0
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				// when grid is island water and not visited
				cnt := islandDfs(grid, visited, i, j)
				if cnt > max {
					max = cnt
				}
			}
		}
	}
	return max
}

func islandDfs(grid [][]int, visited [][]bool, i int, j int) int {
	if i < 0 || i >= len(grid) ||
		j < 0 || j >= len(grid[0]) ||
		visited[i][j] || grid[i][j] == 0 {
		// out of boundary, has visited, water not island
		return 0
	}
	count := 1
	visited[i][j] = true
	for x := range directions {
		newRow := i + directions[x][0]
		newCol := j + directions[x][1]
		count += islandDfs(grid, visited, newRow, newCol)
	}
	return count
}
