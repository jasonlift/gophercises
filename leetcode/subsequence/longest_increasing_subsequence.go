package subsequence

/**
300
https://leetcode.com/problems/longest-increasing-subsequence/
 */

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := make([]int, len(nums))
	for i := range nums {
		memo[i] = 1
	}
	for i:=1; i<len(nums); i++ {
		for j:=0; j<i; j++ {
			if nums[i] > nums[j] {
				memo[i] = max(memo[i], memo[j]+1)
			}
		}
	}

	res := 0
	for i:=0; i<len(nums); i++ {
		res = max(res, memo[i])
	}
	return res
}

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}