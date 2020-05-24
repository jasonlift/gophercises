package array

func quicksort(nums []int, start int, end int) {
	if start < end {
		p := partition(nums, start, end)
		quicksort(nums, start, p)
		quicksort(nums, p+1, end)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	tmp := left
	for left < right {
		left++
		for left < right && pivot > nums[left] {
			left++
		}
		for pivot < nums[right] {
			right--
		}
		if left < right {
			t := nums[left]
			nums[left] = nums[right]
			nums[right] = t
		}
	}
	// left == right
	nums[tmp] = nums[right]
	nums[right] = pivot
	return right
}