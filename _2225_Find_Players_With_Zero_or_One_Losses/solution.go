package _2225_Find_Players_With_Zero_or_One_Losses

import "sort"

func findWinners(matches [][]int) [][]int {
	all := make(map[int]bool)
	lost := make(map[int]int)

	for i := 0; i < len(matches); i++ {
		all[matches[i][0]], all[matches[i][1]] = true, true
		lost[matches[i][1]]++
	}
	result := make([][]int, 2)
	for k, _ := range all {
		if val, exists := lost[k]; exists {
			if val == 1 {
				result[1] = append(result[1], k)
			}
		} else {
			result[0] = append(result[0], k)
		}
	}

	sort.Slice(result[0], func(i, j int) bool {
		return result[0][i] < result[0][j]
	})

	sort.Slice(result[1], func(i, j int) bool {
		return result[1][i] < result[1][j]
	})

	return result
}
