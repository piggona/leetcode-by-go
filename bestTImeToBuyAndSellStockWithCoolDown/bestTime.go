package bestTImeToBuyAndSellStockWithCoolDown

import "math"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy := make([]int, len(prices))
	sell := make([]int, len(prices))
	max := 0
	buy[0] = -prices[0]
	sell[0] = 0
	for i := 1; i < len(prices); i++ {
		delta := prices[i] - prices[i-1]
		sell[i] = getMax(buy[i-1]+prices[i], sell[i-1]+delta)
		if i-2 >= 0 {
			buy[i] = getMax(buy[i-1]-delta, sell[i-2]-prices[i])
		} else {
			buy[i] = buy[i-1] - delta
		}

		if sell[i] > max {
			max = sell[i]
		}
	}
	return max
}

func getMax(nums ...int) int {
	max := math.MinInt32
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}
