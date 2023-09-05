package _643_Maximum_Average_Subarray_I_

func findMaxAverage(nums []int, k int) float64 {
	l, r := 0, 0
	var maxAvg *float64
	curr := 0
	for {
		if r > len(nums)-1 {
			break
		}
		for {
			curr += nums[r]
			r++
			if r-l > k-1 {
				break
			}
		}
		avg := float64(curr) / float64(k)
		if maxAvg == nil || *maxAvg < avg {
			maxAvg = &avg
		}
		curr -= nums[l]
		l++
	}

	return *maxAvg
}
