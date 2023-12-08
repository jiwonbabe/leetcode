package _20_Valid_Parentheses_

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if s[i] == ')' {
				if top == '(' {
					stack = stack[0 : len(stack)-1]
				} else {
					return false
				}
			} else if s[i] == '}' {
				if top == '{' {
					stack = stack[0 : len(stack)-1]
				} else {
					return false
				}
			} else {
				if top == '[' {
					stack = stack[0 : len(stack)-1]
				} else {
					return false
				}
			}
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
