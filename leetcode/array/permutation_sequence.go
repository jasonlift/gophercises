package array

import (
	"strconv"
	"strings"
)

/**
60
https://leetcode.com/problems/permutation-sequence/
 */

func getPermutation(n int, k int) string {
	nums := make([]string, n)
	for i:=1; i<=n; i++ {
		nums[i-1] = strconv.Itoa(i)
	}
	counter := 0
	tmp := []string{}
	haveSeen := make([]bool, n)
	res := seqBacktrack(nums, &tmp, haveSeen, &counter, k)
	return strings.Join(res, "")
}

func getPermutationII(n int, k int) string {
	nums := make([]string, n)
	fact := 1
	for i:=1; i<=n; i++ {
		nums[i-1] = strconv.Itoa(i)
		if i != n {
			fact *= i
		}
	}

	res := []string{}
	k = k -1 // convert the sequence begin at 1 -> at 0
	round := n-1
	for round >= 0 {
		idx := k / fact
		res = append(res, nums[idx])
		k = k % fact
		nums = append(nums[0:idx], nums[idx+1:]...)
		if round > 0 {
			fact = fact / round
		}
		round--
	}
	return strings.Join(res, "")
}

func seqBacktrack(nums []string, tmp *[]string, haveSeen []bool, pcnt *int, k int) []string {
	if len(*tmp) == len(nums) {
		*pcnt++
		if *pcnt == k {
			res := make([]string, len(nums))
			copy(res, *tmp)
			return res
		}
	}

	for i:=0; i<len(nums); i++ {
		if haveSeen[i] {
			continue
		}
		*tmp = append(*tmp, nums[i]) // 和下方对称
		haveSeen[i] = true // 和下方对称
		res := seqBacktrack(nums, tmp, haveSeen, pcnt, k)
		if res != nil {
			return res
		}
		haveSeen[i] = false
		*tmp = (*tmp)[:len(*tmp)-1]
	}
	return nil
}