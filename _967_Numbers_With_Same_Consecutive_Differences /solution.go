package _967_Numbers_With_Same_Consecutive_Differences_

import (
	"strconv"
	"strings"
)

func numsSameConsecDiff(n int, k int) []int {
	outputSet := make(map[int]bool)
	backtracking(n, k, &outputSet, []int{})
	output := make([]int, 0)
	for k, _ := range outputSet {
		output = append(output, k)
	}
	return output
}

func backtracking(n int, k int, output *map[int]bool, curr []int) {
	if len(curr) == n {
		var sb strings.Builder
		for i := 0; i < len(curr); i++ {
			sb.WriteString(strconv.Itoa(curr[i]))
		}
		result, _ := strconv.Atoi(sb.String())
		(*output)[result] = true
		return
	}
	if len(curr) == 0 {
		for i := 1; i < 10; i++ {
			curr = append(curr, i)
			backtracking(n, k, output, curr)
			curr = curr[0 : len(curr)-1]
		}
	} else {
		last := curr[len(curr)-1]
		cand1 := last - k
		if cand1 < 10 && cand1 >= 0 {
			curr = append(curr, cand1)
			backtracking(n, k, output, curr)
			curr = curr[0 : len(curr)-1]
		}

		cand2 := last + k
		if cand2 < 10 && cand2 >= 0 {
			curr = append(curr, cand2)
			backtracking(n, k, output, curr)
			curr = curr[0 : len(curr)-1]
		}
	}
}
