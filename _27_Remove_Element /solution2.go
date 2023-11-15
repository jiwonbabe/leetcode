package _27_Remove_Element_

func removeElement2(nums []int, val int) int {
	r := len(nums) - 1
	l := 0
	cnt := 0
	for l <= r {
		if nums[l] == val {
			nums[l], nums[r] = nums[r], -1
			r--
		} else {
			cnt++
			l++
		}
	}

	return cnt
}
