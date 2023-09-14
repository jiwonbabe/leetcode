package _1189__Maximum_Number_of_Balloons_

import "math"

func maxNumberOfBalloons(text string) int {
	cntArr := make([]int, 5)
	for i := 0; i < len(text); i++ {
		if text[i] == 'b' {
			cntArr[0]++
		} else if text[i] == 'a' {
			cntArr[1]++
		} else if text[i] == 'l' {
			cntArr[2]++
		} else if text[i] == 'o' {
			cntArr[3]++
		} else if text[i] == 'n' {
			cntArr[4]++
		}
	}
	cntArr[2] = cntArr[2] / 2
	cntArr[3] = cntArr[3] / 2
	min := float64(cntArr[0])
	for _, value := range cntArr {
		min = math.Min(min, float64(value))
	}

	return int(min)
}
