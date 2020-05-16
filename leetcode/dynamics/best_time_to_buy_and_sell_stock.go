package dynamics

/**
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

Given an array, find the maximum gap between different index
 */

/**
两个变量存储临时状态
用空间lowestPrice存储了线性的历史价格
不用每次都从开始遍历了
 */
func maxProfit(prices []int) int {
	if prices == nil || len(prices) <=1 {
		return 0
	}

	lowestPrice := prices[0]
	max := 0
	i := 1
	for i < len(prices) {
		if prices[i] < lowestPrice {
			lowestPrice = prices[i]
		}

		currProfit := prices[i] - lowestPrice
		if currProfit > max {
			max = currProfit
		}
		i++
	}
	return max
}


/**
查了一个三角的表，两层循环，效率低
 */
func maxProfitSelf(prices []int) int {
	if prices == nil || len(prices) <=1 {
		return 0
	}
	max := 0
	for i:=0; i<len(prices); i++ {
		for j:=len(prices)-1; j>i; j-- {
			profit := prices[j]-prices[i]
			if profit > max {
				max = profit
			}
		}
	}
	return max
}