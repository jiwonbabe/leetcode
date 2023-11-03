package _22_Generate_Parentheses_

func generateParenthesis(n int) []string {
	results := make([]string, 0)
	backtracking(&results, 0, 0, n, []byte{})
	return results
}

func backtracking(results *[]string, leftCnt int, rightCnt int, n int, comb []byte) {
	if len(comb) == 2*n {
		combCopy := make([]byte, len(comb))
		copy(combCopy, comb)
		*results = append(*results, string(combCopy))
		return
	}
	if leftCnt < n {
		comb = append(comb, '(')
		backtracking(results, leftCnt+1, rightCnt, n, comb)
		comb = comb[0 : len(comb)-1]
	}
	if rightCnt < leftCnt { // rightCnt >= leftCnt 이면 prefix가 ) 인 string 생길 수 있기 때문.
		comb = append(comb, ')')
		backtracking(results, leftCnt, rightCnt+1, n, comb)
		comb = comb[0 : len(comb)-1]
	}
}
