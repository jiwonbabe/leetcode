package _383__Ransom_Note_

func canConstruct(ransomNote string, magazine string) bool {
	mhm := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		mhm[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if v, _ := mhm[ransomNote[i]]; v > 0 {
			mhm[ransomNote[i]]--
		} else {
			return false
		}
	}

	return true
}
