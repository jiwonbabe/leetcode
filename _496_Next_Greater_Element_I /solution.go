package _496_Next_Greater_Element_I_

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0)
	hm := make(map[int]int)
	result := make([]int, 0)

	for _, n := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < n {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			hm[top] = n
		}
		stack = append(stack, n)
	}

	for len(stack) > 0 {
		hm[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	for _, n := range nums1 {
		result = append(result, hm[n])
	}

	return result
}
