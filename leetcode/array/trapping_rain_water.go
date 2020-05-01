package array

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	l, r, lmax, rmax := 0, len(height)-1, height[0], height[len(height)-1]
	var res int = 0
	for l < r {
		if lmax < rmax {
			l++
			if lmax > height[l] {
				res += lmax-height[l]
			} else {
				lmax = height[l]
			}
		} else {
			r--
			if rmax > height[r] {
				res += rmax - height[r]
			} else {
				rmax = height[r]
			}
		}
	}
	return res
}

func trapV2(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}
	return res
}