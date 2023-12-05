package _56_Merge_Intervals_

import "sort"

func merge(intervals [][]int) [][]int {
	output := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var curr []int
	for i := 0; i < len(intervals); i++ {
		if curr == nil {
			curr = intervals[i]
		} else {
			if intervals[i][0] <= curr[1] {
				if curr[1] < intervals[i][1] {
					curr[1] = intervals[i][1]
				}
			} else {
				output = append(output, curr)
				curr = intervals[i]
			}
		}
	}

	output = append(output, curr)

	return output
}
