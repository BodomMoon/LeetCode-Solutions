// Time:  O(n)
// Space: O(1)

func isPalindrome(s string) bool {
	r, _ := regexp.Compile("[^a-zA-Z0-9]")
	s = r.ReplaceAllString(s, "")
	return checkPalindrome(s)
}

func checkPalindrome(s string) bool{
	for i := 0; i < len(s)/2; i++ {
        if (s[i] | 32) != (s[len(s)-1-i] | 32) {
			return false
		}
	}
	return true
}
