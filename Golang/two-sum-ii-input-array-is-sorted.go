// Time:  O(n)
// Space: O(1)

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for sum := numbers[left] + numbers[right]; left < right; sum = numbers[left] + numbers[right] {
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			break
		}
	}
	return []int{left + 1, right + 1}
}