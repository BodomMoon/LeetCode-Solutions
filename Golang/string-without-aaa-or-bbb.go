// Time:  O(a + b)
// Space: O(1)

type Pair struct {
	count, letter interface{}
}

func strWithout3a3b(A int, B int) string {
	buffer := new(bytes.Buffer)
	max, min := getMaxPair(&A, &B)
	for i, maxCount, length := 0, 0, A+B; i < length; i++ {
		if *min.count.(*int) >= *max.count.(*int) || maxCount == 2 {
			buffer.WriteByte(min.letter.(byte))
			*min.count.(*int) -= 1
			maxCount = 0
		} else {
			buffer.WriteByte(max.letter.(byte))
			*max.count.(*int) -= 1
			maxCount++
		}
	}
	return buffer.String()
}
func getMaxPair(X *int, Y *int) (Pair, Pair) {
	if *X >= *Y {
		return Pair{X, byte('a')}, Pair{Y, byte('b')}
	}
	return Pair{Y, byte('b')}, Pair{X, byte('a')}
}