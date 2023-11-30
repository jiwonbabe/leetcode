package _1_Two_Sum_

func twoSum(nums []int, target int) []int {
	ht := make(map[int]int)
	for i, n := range nums {
		complement := target - n
		if v, e := ht[complement]; e && v != i {
			return []int{i, v}
		}
		ht[n] = i
	}
	return []int{}
}
