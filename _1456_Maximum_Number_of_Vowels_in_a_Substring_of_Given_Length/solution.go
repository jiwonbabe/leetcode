package _1456_Maximum_Number_of_Vowels_in_a_Substring_of_Given_Length

func maxVowels(s string, k int) int {
	i, l, cnt := 0, 0, 0
	maxCnt := 0

	for {
		if i > len(s)-1 {
			break
		}
		for i-l < k {
			if isVowel(s[i]) {
				cnt++
			}
			i++
		}
		if maxCnt < cnt {
			maxCnt = cnt
		}
		if isVowel(s[l]) {
			cnt--
		}
		l++
	}
	return maxCnt
}

func isVowel(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	return false
}
