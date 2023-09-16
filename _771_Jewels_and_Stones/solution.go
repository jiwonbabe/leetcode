package _771_Jewels_and_Stones

func numJewelsInStones(jewels string, stones string) int {
	hm := make(map[byte]int)
	for i := 0; i < len(stones); i++ {
		hm[stones[i]]++
	}
	cnt := 0
	for i := 0; i < len(jewels); i++ {
		cnt += hm[jewels[i]]
	}

	return cnt
}
