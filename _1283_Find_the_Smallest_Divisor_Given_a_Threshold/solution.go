package _1283_Find_the_Smallest_Divisor_Given_a_Threshold

import "math"

func smallestDivisor(nums []int, threshold int) int {
	l, r := 1, nums[0]
	for _, n := range nums {
		if r < n {
			r = n
		}
	}
	ans := 0
	for l <= r {
		mid := l + (r-l)/2
		if check(nums, mid, threshold) {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}

func check(nums []int, div int, threshold int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += int(math.Ceil(float64(nums[i]) / float64(div)))
		if sum > threshold {
			return false
		}
	}
	return true
}
