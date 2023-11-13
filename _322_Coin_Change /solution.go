package _322_Coin_Change_

import "math"

func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -1
	}
	memo[0] = 0

	res := dp(coins, amount, memo)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func dp(coins []int, amount int, memo []int) int {
	if amount < 0 {
		return math.MaxInt32
	}
	if memo[amount] != -1 {
		return memo[amount]
	}

	min := math.MaxInt32
	for _, coin := range coins {
		res := dp(coins, amount-coin, memo)
		if res != math.MaxInt32 && res+1 < min {
			min = res + 1
		}
	}

	memo[amount] = min
	return memo[amount]
}
