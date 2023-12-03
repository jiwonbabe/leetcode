package _128_Longest_Consecutive_Sequence_

func longestConsecutive(nums []int) int {
	hs := make(map[int]bool)
	for _, n := range nums {
		hs[n] = true
	}

	longest := 0
	for k, _ := range hs {
		if _, e := hs[k-1]; !e {
			currLen := 0
			currVal := k
			for {
				if _, e := hs[currVal]; !e {
					break
				}
				currLen++
				currVal++
			}
			if longest < currLen {
				longest = currLen
			}
		}
	}

	return longest
}
