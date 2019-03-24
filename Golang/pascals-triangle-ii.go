// Time:  O(n^2)
// Space: O(1)

func getRow(rowIndex int) []int {
	var result []int
	for i := 0; i <= rowIndex; i++ {
		for j, tmp := 1, 1; j < i; j++ {
			result[j], tmp = tmp, result[j]
			result[j] += tmp
		}
		result = append(result, 1)
	}
	return result
}