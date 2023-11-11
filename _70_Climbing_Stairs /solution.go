package _70_Climbing_Stairs_

func climbStairsRecursive(n int) int {
	memo := make([]int, n+1)
	return dp(n, memo)
}

func dp(i int, memo []int) int {
	if i == 1 {
		memo[1] = 1
		return 1
	}
	if i == 2 {
		memo[1] = 2
		return 2
	}
	if memo[i] != 0 {
		return memo[i]
	}
	memo[i] = dp(i-1, memo) + dp(i-2, memo)
	return memo[i]
}

func climbStairsIterative(n int) int {
	if n <= 2 {
		return n
	}
	s := 1
	b := 2
	dp := 0
	for i := 3; i < n+1; i++ {
		dp = s + b
		s = b
		b = dp
	}
	return dp
}
