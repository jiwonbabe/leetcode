package _15_3Sum_

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	output := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i == 0 || nums[i-1] != nums[i] {
			l, h := i+1, len(nums)-1
			for l < h {
				sum := nums[i] + nums[l] + nums[h]
				if sum < 0 {
					l++
				} else if sum > 0 {
					h--
				} else {
					output = append(output, []int{nums[i], nums[l], nums[h]})
					for l < h && nums[l] == nums[l+1] {
						l++ // Skip duplicate
					}
					for l < h && nums[h] == nums[h-1] {
						h-- // Skip duplicate
					}
					l++
					h--
				}
			}
		}
	}
	return output
}
