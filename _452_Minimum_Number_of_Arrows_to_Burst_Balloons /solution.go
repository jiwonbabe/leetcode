package _452_Minimum_Number_of_Arrows_to_Burst_Balloons_

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	cnt := 0
	overlapped := points[0]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= overlapped[1] {
			overlapped = merge(overlapped, points[i])
		} else {
			cnt++
			overlapped = points[i]
		}
	}

	cnt++

	return cnt
}

func merge(i1, i2 []int) []int {
	return []int{i2[0], min(i1[1], i2[1])}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
