package _746__Min_Cost_Climbing_Stairs_

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	memo := make([]int, n+1)
	return dp(cost, memo, n)
}

func dp(cost []int, memo []int, i int) int {
	if i <= 1 {
		memo[i] = 0
		return memo[i]
	}
	if memo[i] != 0 {
		return memo[i]
	}
	twoStep := dp(cost, memo, i-2) + cost[i-2]
	oneStep := dp(cost, memo, i-1) + cost[i-1]
	if oneStep < twoStep {
		memo[i] = oneStep
	} else {
		memo[i] = twoStep
	}
	return memo[i]
}

func minCostClimbingStairsIterative(cost []int) int {
	n := len(cost)
	oneStep, twoStep, curr := 0, 0, 0
	for i := 2; i < n+1; i++ {
		if oneStep+cost[i-1] < twoStep+cost[i-2] {
			curr = oneStep + cost[i-1]
		} else {
			curr = twoStep + cost[i-2]
		}
		twoStep = oneStep
		oneStep = curr
	}

	return curr
}
