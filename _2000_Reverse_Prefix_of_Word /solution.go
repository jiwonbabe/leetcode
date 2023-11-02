package _2000_Reverse_Prefix_of_Word_

func reversePrefix(word string, ch byte) string {
	l := 0
	wordBytes := []byte(word)
	for i := 0; i < len(word); i++ {
		if wordBytes[i] == ch {
			for l <= i {
				wordBytes[l], wordBytes[i] = wordBytes[i], wordBytes[l]
				l++
				i--
			}
			break
		}
	}

	return string(wordBytes)
}
