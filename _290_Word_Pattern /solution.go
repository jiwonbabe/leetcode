package _290_Word_Pattern_

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	ht := make(map[string]int)
	index := 0
	for i := 0; i < len(pattern); i++ {
		converted := string(pattern[i]) + " "
		if _, e := ht[converted]; !e {
			ht[converted] = index
		}
		if _, e := ht[words[i]]; !e {
			ht[words[i]] = index
		}
		if ht[converted] != ht[words[i]] {
			return false
		}
		index++
	}
	return true
}
