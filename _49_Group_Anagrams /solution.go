package _49_Group_Anagrams_

import (
	"fmt"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	output := make([][]string, 0)
	ht := make(map[string][]string)
	for _, str := range strs {
		count := make([]int, 26)
		for i := 0; i < len(str); i++ {
			count[int(str[i])-int('a')]++
		}
		var sb strings.Builder
		for _, v := range count {
			sb.WriteString(fmt.Sprintf("%d,", v))
		}
		key := sb.String()
		if _, e := ht[key]; e {
			ht[key] = append(ht[key], str)
		} else {
			ht[key] = []string{str}
		}
	}

	for _, v := range ht {
		output = append(output, v)
	}
	return output
}
