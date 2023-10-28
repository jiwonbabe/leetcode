package _2389__Longest_Subsequence_With_Limited_Sum

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	sum := 0
	for i, _ := range nums {
		sum += nums[i]
		nums[i] = sum
	}

	for i, q := range queries {
		l, r := 0, len(nums)-1
		for l <= r {
			mid := l + (r-l)/2
			if nums[mid] == q {
				queries[i] = mid + 1
				break
			} else if nums[mid] < q {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if l > r {
			queries[i] = l
		}
	}

	return queries
}
