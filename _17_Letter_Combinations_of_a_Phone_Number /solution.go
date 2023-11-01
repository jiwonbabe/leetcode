package _17_Letter_Combinations_of_a_Phone_Number_

func letterCombinations(digits string) []string {
	ht := make(map[byte][]byte)
	ht['2'] = []byte{'a', 'b', 'c'}
	ht['3'] = []byte{'d', 'e', 'f'}
	ht['4'] = []byte{'g', 'h', 'i'}
	ht['5'] = []byte{'j', 'k', 'l'}
	ht['6'] = []byte{'m', 'n', 'o'}
	ht['7'] = []byte{'p', 'q', 'r', 's'}
	ht['8'] = []byte{'t', 'u', 'v'}
	ht['9'] = []byte{'w', 'x', 'y', 'z'}

	results := make([]string, 0)

	if len(digits) > 0 {
		backtracking(digits, &results, []byte{}, ht)
	}
	return results
}

func backtracking(digits string, results *[]string, comb []byte, ht map[byte][]byte) {
	if len(comb) == len(digits) {
		*results = append(*results, string(comb))
		return
	}
	for _, c := range ht[digits[len(comb)]] {
		comb = append(comb, c)
		backtracking(digits, results, comb, ht)
		comb = comb[:len(comb)-1]
	}
}
