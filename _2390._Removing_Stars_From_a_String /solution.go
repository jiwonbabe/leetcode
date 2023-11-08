package _2390__Removing_Stars_From_a_String_

func removeStars(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
