package dynamics

/**
https://leetcode.com/problems/maximum-subarray/
 */

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	sum := nums[0]
	for _, i := range nums[1:] {
		if sum > 0 {
			sum += i
		} else {
			sum = i
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxSubArraySelf(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	var sum int = 0
	for i:=0; i<len(nums); i++ {
		t := sum + nums[i]
		result = max(result, t)
		if t > 0 {
			sum = t
		} else {
			sum = 0
		}
	}
	return result
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}