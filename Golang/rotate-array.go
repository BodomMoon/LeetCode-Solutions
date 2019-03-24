// Time:  O(n)
// Space: O(1)
func rotate(nums []int, k int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}

	for i := 0; i < k%len(nums)/2; i++ {
		nums[i], nums[k%len(nums)-1-i] = nums[k%len(nums)-1-i], nums[i]
	}

	for i := 0; i < (len(nums)-k%len(nums))/2; i++ {
		nums[i+k%len(nums)], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i+k%len(nums)]
	}

	return
}

// Time:  O(n)
// Space: O(n)
func rotate2(nums []int, k int) {
	copy(nums, append(nums[len(nums)-k%len(nums):], nums[:len(nums)-k%len(nums)]...))
	return
}