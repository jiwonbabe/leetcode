package _121_Best_Time_to_Buy_and_Sell_Stock_

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0
	for i := 0; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
		cand := prices[i] - minPrice
		if profit < cand {
			profit = cand
		}
	}

	return profit
}
