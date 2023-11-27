package _205_Isomorphic_Strings_

func isIsomorphic(s string, t string) bool {
	st := make(map[byte]byte)
	tt := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, e := st[s[i]]; e {
			if v != t[i] {
				return false
			}
		} else {
			st[s[i]] = t[i]
		}

		if v, e := tt[t[i]]; e {
			if v != s[i] {
				return false
			}
		} else {
			tt[t[i]] = s[i]
		}
	}
	return true
}
