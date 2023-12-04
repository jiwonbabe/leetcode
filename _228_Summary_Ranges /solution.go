package _228_Summary_Ranges_

import "fmt"

func summaryRanges(nums []int) []string {
	ranges := make([]string, 0)
	i := 0

	for i < len(nums) {
		start := nums[i]
		for i+1 < len(nums) && nums[i]+1 == nums[i+1] {
			i++
		}

		if start != nums[i] {
			ranges = append(ranges, fmt.Sprintf("%d->%d", start, nums[i]))
		} else {
			ranges = append(ranges, fmt.Sprintf("%d", start))
		}
		i++
	}

	return ranges
}
