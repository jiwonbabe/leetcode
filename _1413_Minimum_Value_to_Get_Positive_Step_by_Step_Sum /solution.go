package _1413_Minimum_Value_to_Get_Positive_Step_by_Step_Sum_

func minStartValue(nums []int) int {
	min := 100*len(nums) + 1
	ps := 0
	for i := 0; i < len(nums); i++ {
		ps += nums[i]
		nums[i] = ps
		if min > ps {
			min = ps
		}
	}

	if min >= 0 {
		return 1
	}
	return -1*min + 1
}
