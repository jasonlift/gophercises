package array

/**
https://leetcode.com/problems/friend-circles/

 */

func findCircleNum(M [][]int) int {
	count := 0
	visited := make([]int, len(M))

	for i:=0; i<len(M); i++ {
		if visited[i] == 0 {
			count++
			friendDfs(M, visited, i)
		}
	}

	return count
}

func friendDfs(M [][]int, visited []int, i int) {
	for j:=0; j<len(M); j++ {
		if M[i][j] == 1 && visited[j] == 0 {
			visited[j] = 1
			friendDfs(M, visited, j)
		}
	}
}
