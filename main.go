package main

func longestPalindrome(s string) string {
	count := 0
	max := 0
	for i := 0; i < len(s); i++ {
		l := i
		for j := len(s) - 1; j > -1; j-- {
			if count == (j - l) {
				return s[i : i-j]
			}
			if s[j] == s[l] {
				l += 1
				count += 2
			} else {
				l = i
				count = 0
			}
		}
	}
	return ""
}

func main() {
	s := "babad"
	longestPalindrome(s)
}
