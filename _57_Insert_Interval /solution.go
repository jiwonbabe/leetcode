package _57_Insert_Interval_

func insert(intervals [][]int, newInterval []int) [][]int {
	output := make([][]int, 0)
	index := 0
	for index < len(intervals) && intervals[index][1] < newInterval[0] {
		output = append(output, intervals[index])
		index++
	}

	for index < len(intervals) && intervals[index][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[index][0])
		newInterval[1] = max(newInterval[1], intervals[index][1])
		index++
	}

	output = append(output, newInterval)

	for i := index; i < len(intervals); i++ {
		output = append(output, intervals[i])
	}

	return output

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
