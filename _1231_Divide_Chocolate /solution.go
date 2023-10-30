package _1231_Divide_Chocolate_

func maximizeSweetness(sweetness []int, k int) int {
	var l, r int
	l = 1
	sum := 0
	for i, _ := range sweetness {
		sum += sweetness[i]
		if l > sweetness[i] {
			l = sweetness[i]
		}
	}
	r = sum
	for l <= r {
		mid := l + (r-l+1)/2
		if check(mid, k, sweetness) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l - 1
}

func check(mid int, k int, sweetness []int) bool {
	cnt := 0
	currSum := 0
	for i, _ := range sweetness {
		currSum += sweetness[i]
		if currSum >= mid {
			cnt++
			currSum = 0
		}
	}
	if cnt >= k+1 {
		return true
	}
	return false
}
