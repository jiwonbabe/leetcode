package _80_Remove_Duplicates_from_Sorted_Array_II_

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	i := 2
	for j := 2; j < len(nums); j++ {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
