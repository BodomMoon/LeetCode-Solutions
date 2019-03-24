// Time:  O(n)
// Space: O(n)

func countPrimes(n int) int {
    if (n <= 2) {
        return 0
    }

    isPrimes := make([]bool, n)
    result := n/2
    for i := 3; i * i < n; i += 2 {
        if !isPrimes[i] == false {
            continue
        }
        for j := i * i ; j < n; j += i * 2 {
            if !isPrimes[j] == false {
                continue
            }
            result--
            isPrimes[j] = true
        }
    }
    return result
}