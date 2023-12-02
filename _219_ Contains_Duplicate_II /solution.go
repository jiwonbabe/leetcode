package _219__Contains_Duplicate_II_

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	ht := make(map[int][]int)
	for i, n := range nums {
		ht[n] = append(ht[n], i)
	}

	for _, indexes := range ht {
		if len(indexes) < 2 {
			continue
		}
		for i := 0; i < len(indexes)-1; i++ {
			if indexes[i+1]-indexes[i] <= k {
				return true
			}
		}
	}

	return false
}
