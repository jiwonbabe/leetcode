package _3__Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring(s string) int {
	max := 0
	l := 0
	h := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if _, e := h[s[i]]; e {
			if max < i-l {
				max = i - l
			}
			for h[s[i]] {
				delete(h, s[l])
				l++
			}
		}
		h[s[i]] = true
	}

	if len(s)-l > max {
		max = len(s) - l
	}

	return max
}
