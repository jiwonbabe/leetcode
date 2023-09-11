package _1133_Largest_Unique_Number

func largestUniqueNumber(nums []int) int {
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]]++
	}
	r := -1
	for k, v := range hashmap {
		if v == 1 && r < k {
			r = k
		}
	}
	return r
}
