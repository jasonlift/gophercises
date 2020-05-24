package array

/*
Description:
https://leetcode.com/problems/kth-largest-element-in-an-array/

Example 1:
Input: [3,2,1,5,6,4] and k = 2
Output: 5

Example 2:
Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
 */

// using quicksort method
func findKthLargest(nums []int, k int) int {
	return kQuicksort(&nums, 0, len(nums)-1, k)
}

func kQuicksort(nums *[]int, l int, r int, k int) int {
	if l < r {
		p := kPartition(nums, l, r)
		if p == k-1 {
			return (*nums)[p]
		}

		if p >= k {
			kQuicksort(nums, l, p-1, k)
		} else {
			kQuicksort(nums, p+1, r, k)
		}
	}
	return (*nums)[k-1] // boundary?
}

func kPartition(nums *[]int, l, r int) int {
	pivot := (*nums)[l]
	tmp := l
	for l < r {
		l++
		for l <= r && pivot <= (*nums)[l] { //inverted order
			l++
		}
		for pivot > (*nums)[r] {
			r--
		}

		if l < r {
			//swap
			t := (*nums)[l]
			(*nums)[l] = (*nums)[r]
			(*nums)[r] = t
		}
	}
	// this moment, l == r
	(*nums)[tmp] = (*nums)[r]
	(*nums)[r] = pivot
	return r
}