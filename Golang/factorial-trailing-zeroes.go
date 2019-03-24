// Time:  O(logn) = 1
// Space: O(1)

func trailingZeroes(n int) int {
	result := 0
	for n != 0 {
		n /= 5
		result += n
	}
	return result
}