//Time:  O(k) where k is the steps to be happy number
//Space: O(1)

func isHappy(n int) bool {
    const maxSUm = 9 * 9 * 10
    var resultHash [maxSUm]bool
    sum := 0
    num := n
    for sum != 1 {
        strForm := strconv.Itoa(num)
        byteForm := []byte(strForm)
        sum = 0
        for index := range byteForm {
            sum += int(byteForm[index] - '0')*int(byteForm[index] - '0')
        }
        if resultHash[sum] == true {
            return false
        }
        resultHash[sum] = true
        num = sum
    }
    return true
}