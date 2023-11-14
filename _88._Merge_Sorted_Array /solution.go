package _88__Merge_Sorted_Array_

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Make a copy of the first m elements of nums1
	nums1Copy := make([]int, m)
	copy(nums1Copy, nums1[:m])

	// Read pointers for nums1Copy and nums2 respectively
	p1, p2 := 0, 0

	// Compare elements from nums1Copy and nums2 and write the smallest to nums1
	for p := 0; p < n+m; p++ {
		// Ensure that p1 and p2 aren't over the boundaries of their respective arrays
		if p2 >= n || (p1 < m && nums1Copy[p1] <= nums2[p2]) {
			nums1[p] = nums1Copy[p1]
			p1++
		} else {
			nums1[p] = nums2[p2]
			p2++
		}
	}
}
