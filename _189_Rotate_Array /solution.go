package _189_Rotate_Array_

func rotate(nums []int, k int) {
	n := len(nums)
	if n < 2 {
		return
	}

	// Normalize k to be within the bounds of nums length
	k = k % n
	if k == 0 {
		return // No rotation needed
	}

	// Reverse the first part of the array
	reverse(nums, 0, n-k-1)

	// Reverse the second part of the array
	reverse(nums, n-k, n-1)

	// Reverse the entire array
	reverse(nums, 0, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
