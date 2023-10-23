package _1323_Maximum_69_Number_

import "strconv"

func maximum69Number(num int) int {
	origin := strconv.Itoa(num)
	changedArr := make([]byte, len(origin))
	isChanged := false
	for i, d := range origin {
		if d == '6' && !isChanged {
			changedArr[i] = '9'
			isChanged = true
		} else {
			changedArr[i] = byte(d)
		}
	}

	result, _ := strconv.Atoi(string(changedArr))
	return result
}
