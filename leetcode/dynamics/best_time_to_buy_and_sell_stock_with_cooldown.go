package dynamics

/**
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
 */

func maxProfitII(prices []int) int {
	if prices == nil || len(prices) <=1 {
		return 0
	}

	profit := 0
	lowPrice := prices[0]
	i := 1
	for i < len(prices) {
		if lowPrice > prices[i] {
			 lowPrice = prices[i]
		}

		if prices[i] - lowPrice > 0 {
			profit += prices[i] - lowPrice
			lowPrice = prices[i]
		}
		i++
	}
	return profit
}