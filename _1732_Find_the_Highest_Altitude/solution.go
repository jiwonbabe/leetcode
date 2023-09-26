package _1732_Find_the_Highest_Altitude

func largestAltitude(gain []int) int {
	max := 0
	curr := 0
	for _, g := range gain {
		curr = curr + g
		if curr > max {
			max = curr
		}
	}

	return max
}
