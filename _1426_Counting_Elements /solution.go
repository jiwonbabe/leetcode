package _1426_Counting_Elements_

func countElements(arr []int) int {
	hm := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, exists := hm[arr[i]]; exists {
			hm[arr[i]] += 1
		} else {
			hm[arr[i]] = 1
		}
	}
	cnt := 0
	for k, v := range hm {
		if _, exists := hm[k+1]; exists {
			cnt += v
		}
	}

	return cnt
}
