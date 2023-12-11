package _150_Evaluate_Reverse_Polish_Notation

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			result := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		} else if tokens[i] == "-" {
			result := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		} else if tokens[i] == "*" {
			result := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		} else if tokens[i] == "/" {
			result := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		} else {
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}
