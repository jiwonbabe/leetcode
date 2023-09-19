package _1544_Make_The_String_Great_

func makeGood(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			diff := int(top) - int(s[i])
			if diff == 32 || diff == -32 {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else {
			stack = append(stack, s[i])
		}

	}

	return string(stack)
}
