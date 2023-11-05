package _216_Combination_Sum_III_

func combinationSum3(k int, n int) [][]int {
	output := make([][]int, 0)
	backtracking(k, n, &output, []int{}, 0)
	return output
}

func backtracking(k int, n int, output *[][]int, curr []int, sum int) {
	if len(curr) == k {
		if sum == n {
			currCopy := make([]int, len(curr))
			copy(currCopy, curr)
			*output = append(*output, currCopy)
			return
		}
		return
	}
	if len(curr) == 0 {
		for i := 1; i < 10; i++ {
			curr = append(curr, i)
			backtracking(k, n, output, curr, i)
			curr = curr[0 : len(curr)-1]
		}
	} else {
		last := curr[len(curr)-1]
		for i := last + 1; i < 10; i++ {
			curr = append(curr, i)
			backtracking(k, n, output, curr, sum+i)
			curr = curr[0 : len(curr)-1]
		}
	}
}
