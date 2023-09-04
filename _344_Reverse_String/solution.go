package _344_Reverse_String

func reverseString(s []byte) {
	b, e := 0, len(s)-1

	for {
		if b >= e {
			break
		}
		s[b], s[e] = s[e], s[b]
		b++
		e--
	}
}
