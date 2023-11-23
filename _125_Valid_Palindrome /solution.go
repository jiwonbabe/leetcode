package _125_Valid_Palindrome_

import "unicode"

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	l, r := 0, len(s)-1

	for l <= r {
		for !isLetter(rune(s[l])) && l < len(s)-1 {
			l++
		}
		for !isLetter(rune(s[r])) && r > 0 {
			r--
		}
		if l <= r {
			if !isSameLetterIgnoreCase(rune(s[l]), rune(s[r])) {
				return false
			} else {
				l++
				r--
			}
		}
	}
	return true
}

func isLetter(r rune) bool {
	return (unicode.IsNumber(r) || unicode.IsLetter(r)) && !unicode.IsSpace(r)
}

func isSameLetterIgnoreCase(r1, r2 rune) bool {
	return unicode.ToLower(r1) == unicode.ToLower(r2)
}
