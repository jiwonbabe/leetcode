package _283_Move_Zeroes_

func moveZeroes(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		for nums[l] != 0 && l < len(nums)-1 {
			l++
		}
		for nums[r] == 0 && r > 0 {
			r--
		}
		if l >= r {
			break
		}
		for p := l; p < r; p++ {
			nums[p], nums[p+1] = nums[p+1], nums[p]
		}
		r--
	}
}
