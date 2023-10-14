package _917_Reverse_Only_Letters_

import "unicode"

func reverseOnlyLetters(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		if !unicode.IsLetter(runes[i]) {
			i++
		} else if !unicode.IsLetter(runes[j]) {
			j--
		} else {
			runes[i], runes[j] = runes[j], runes[i]
			i++
			j--
		}
	}
	return string(runes)
}
