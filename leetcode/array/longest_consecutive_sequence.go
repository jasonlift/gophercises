package array

/**
128
https://leetcode.com/problems/longest-consecutive-sequence/
 */

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := 1
	dict := make(map[int]int)
	for i := range nums {
		if _, ok := dict[nums[i]]; !ok {
			dict[nums[i]] = 1
			if _, ok := dict[nums[i]-1]; ok {
				l := merge(dict, nums[i]-1, nums[i])
				if l > max {
					max = l
				}
			}
			if _, ok := dict[nums[i]+1]; ok {
				l := merge(dict, nums[i], nums[i]+1)
				if l > max {
					max = l
				}
			}
		}
	}
	return max
}

func merge(dict map[int]int, less int, more int) int {
	left := less-dict[less]+1
	right := more+dict[more]-1
	l := right-left+1
	dict[left] = l
	dict[right] = l
	return l
}