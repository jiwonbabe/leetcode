package _238_Product_of_Array_Except_Self

/*
return answer
answer[i] = product of all elements except nums[i]
product of any prefix or suffix of nums fit in 32-bit int
O(n) time complexity, not using division op
len(answer) = len(nums)
*/
func productExceptSelf(nums []int) []int {
	length := len(nums)
	prefix := make([]int, length)
	suffix := make([]int, length)
	answer := make([]int, length)
	curr := 1
	for i, e := range nums {
		curr *= e
		prefix[i] = curr
	}

	curr = 1
	for i := length - 1; i >= 0; i-- {
		curr *= nums[i]
		suffix[length-1-i] = curr
	}

	for i := 0; i < length; i++ {
		if i == 0 {
			answer[i] = suffix[length-2-i]
		} else if i == length-1 {
			answer[i] = prefix[i-1]
		} else {
			answer[i] = prefix[i-1] * suffix[length-2-i]
		}
	}

	return answer
}
