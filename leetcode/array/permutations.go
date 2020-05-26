package array

/**
46
https://leetcode.com/problems/permutations/
 */

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	helper(&res, nums, 0)
	return res
}

func helper(res *[][]int, nums []int, i int) {
	if i == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for idx:=i; idx<len(nums); idx++ {
		swap(nums, i, idx)
		helper(res, nums, i+1)
		swap(nums, i, idx)
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}