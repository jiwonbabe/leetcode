package _1832_Check_if_the_Sentence_Is_Pangram_

func checkIfPangram(sentence string) bool {
	hashmap := make(map[byte]bool)
	for i := 0; i < len(sentence); i++ {
		hashmap[sentence[i]] = true
	}

	if len(hashmap) == 26 {
		return true
	}
	return false
}
