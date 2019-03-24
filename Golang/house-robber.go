// Time:  O(n)
// Space: O(1)

func rob(nums []int) int {
	result := 0
	last := 0
	for i := range nums {
		last, result = result, max(last+nums[i], result)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}