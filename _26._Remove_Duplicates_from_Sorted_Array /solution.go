package _26__Remove_Duplicates_from_Sorted_Array_

func removeDuplicates(nums []int) int {
	l := 0
	for r := 1; r < len(nums); r++ {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}
	}
	return l + 1
}
