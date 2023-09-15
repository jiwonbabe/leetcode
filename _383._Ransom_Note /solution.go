package _383__Ransom_Note_

func canConstruct(ransomNote string, magazine string) bool {
	rhm := make(map[byte]int)
	mhm := make(map[byte]int)

	for i := 0; i < len(ransomNote); i++ {
		rhm[ransomNote[i]]++
	}
	for i := 0; i < len(magazine); i++ {
		mhm[magazine[i]]++
	}

	for k, v := range rhm {
		if mhm[k] >= v {
			mhm[k] -= v
		} else {
			return false
		}
	}

	return true
}
