// Time:  O(logn)
// Space: O(1)

func convertToTitle(n int) string {
	result := []string{}
	for n != 0 {
		result = append([]string{string('A' + (n-1)%26)}, result...)
		n = (n - 1) / 26
	}
	return strings.Join(result, "")
}