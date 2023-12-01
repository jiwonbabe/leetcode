package _202_Happy_Number

func get_next(n int) int {
	totalSum := 0
	for n > 0 {
		digit := n % 10
		n /= 10
		totalSum += digit * digit
	}
	return totalSum
}

// isHappy determines if a number is a "happy number"
func isHappy(n int) bool {
	seen := make(map[int]struct{})
	for n != 1 {
		if _, exists := seen[n]; exists {
			return false
		}
		seen[n] = struct{}{}
		n = get_next(n)
	}
	return true
}
