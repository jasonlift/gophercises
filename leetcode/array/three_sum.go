package array

import "sort"

/**
15
https://leetcode.com/problems/3sum/

 */

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i:=0; i<len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[i],nums[left],nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left<right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}