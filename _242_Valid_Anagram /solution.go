package _242_Valid_Anagram_

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ht := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		ht[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if _, e := ht[t[i]]; !e {
			return false
		}
		ht[t[i]]--
		if ht[t[i]] == 0 {
			delete(ht, t[i])
		}
	}

	if len(ht) == 0 {
		return true
	}
	return false
}
