package array

/**
33
https://leetcode.com/problems/search-in-rotated-sorted-array/
思路：
二分查找+确定单调区间
 */

func search(nums []int, target int) int {
	nlen := len(nums)
	if nlen == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := (left+right)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[right] { // right part is not rotated
			if nums[mid] > target {
				right = mid-1
				continue
			}

			if target <= nums[right] {
				left = mid+1
			} else {
				right = mid-1
			}
		} else { // left part is not rotated
			if nums[mid] < target {
				left = mid+1
				continue
			}

			if target >= nums[left] {
				right = mid-1
			} else {
				left = mid+1
			}
		}
	}

	if left == right && nums[left] == target {
		return left
	}
	return -1
}
