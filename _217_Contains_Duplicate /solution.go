package _217_Contains_Duplicate_

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = true
	}
	return false
}
