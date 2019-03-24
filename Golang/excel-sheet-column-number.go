// Time:  O(n)
// Space: O(1)

func titleToNumber(s string) int {
	result := 0
	for _, c := range s {
		result = result*26 + int(c) - 'A' + 1
	}
	return result
}