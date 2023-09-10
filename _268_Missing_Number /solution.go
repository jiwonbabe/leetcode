package _268_Missing_Number_

func missingNumber(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	for i := 0; i < len(nums)+1; i++ {
		if _, exists := m[i]; !exists {
			return i
		}
	}
	return -1
}
