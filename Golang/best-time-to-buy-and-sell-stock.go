// Time:  O(n)
// Space: O(1)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	const MaxInt = int((^uint(0)) >> 1)
	minPrice, result := MaxInt, MaxInt
	for _, price := range prices {
		minPrice = min(minPrice, price)
		result = min(result, minPrice-price)
	}
	return -result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}