package _53_Maximum_Subarray

/*
Find the subarray with the largest sum
Sliding window technique
*/
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := 0
	for _, n := range nums {
		if currSum < 0 {
			currSum = 0
		}
		currSum += n
		if maxSum < currSum {
			maxSum = currSum
		}
	}
	return maxSum
}
