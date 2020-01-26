package bestTimeToBuyAndSellStock

func maxProfit(prices []int) int {
	prices = append(prices, 0)
	stack := []int{}
	pos := 0
	res := 0
	for i := 0; i < len(prices); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			pos = i
		} else {
			top := stack[len(stack)-1]
			if prices[i] > prices[top] {
				stack = append(stack, i)
			} else {
				stack = stack[:(len(stack) - 1)]
				temp := prices[top] - prices[pos]
				if temp > res {
					res = temp
				}
				i--
			}
		}
	}
	return res
}
