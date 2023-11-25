package _11_Container_With_Most_Water_

import "math"

func maxArea(height []int) int {
	area := 0
	l, r := 0, len(height)-1
	for l < r {
		cand := (r - l) * int(math.Min(float64(height[l]), float64(height[r])))
		if area < cand {
			area = cand
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return area
}
