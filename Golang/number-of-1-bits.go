//Time:  O(logn)
//Space: O(1)

func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		result += int(num & 1)
		num >>= 1
	}
	return result
}