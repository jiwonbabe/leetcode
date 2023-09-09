package _209_Minimum_Size_Subarray_Sum_

func minSubArrayLen(target int, nums []int) int {
	l, curr := 0, 0
	minSize := 100001
	for i := 0; i < len(nums); i++ {
		curr += nums[i]
		for curr >= target {
			length := i - l + 1
			if minSize > length {
				minSize = length
			}
			curr -= nums[l]
			l++
		}
	}

	if minSize == 100001 {
		return 0
	} else {
		return minSize
	}
}
