package _977_Squares_of_a_Sorted_Array

// 뒤에서부터 채우는 이유는 Array 의 각 원소들을 square 하면 the most biggest value 가 새로 생기는 것이지 the most smallest 가 생기는것은 아니기 때문.
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	l, r := 0, len(nums)-1
	for i := len(result) - 1; i > -1; i-- {
		if absInt(nums[l]) > absInt(nums[r]) {
			result[i] = nums[l] * nums[l]
			l++
		} else {
			result[i] = nums[r] * nums[r]
			r--
		}
	}

	return result
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
