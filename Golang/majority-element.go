// Time:  O(n)
// Space: O(1)

func majorityElement(nums []int) int {
	result, cnt := nums[0], 0
	for _, num := range nums {
		if num == result {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				result, cnt = num, 1
			}
		}
	}
	return result
}