// Time:  O(n^2)
// Space: O(1)

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i] = append(result[i], 1)
			} else {
				result[i] = append(result[i], result[i-1][j-1]+result[i-1][j])
			}
		}
	}
	return result
}