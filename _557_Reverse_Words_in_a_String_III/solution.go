package _557_Reverse_Words_in_a_String_III

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for i := 0; i < len(words); i++ {
		runes := []rune(words[i])
		l, r := 0, len(words[i])-1
		for l <= r {
			runes[l], runes[r] = runes[r], runes[l]
			l++
			r--
		}
		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}
