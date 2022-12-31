package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	longestStr := string(s[0])
	for l := 0; l < len(s)-1; l++ {
		i := l
		r := len(s) - 1
		if (r - l + 1) < len(longestStr) {
			break
		}
		for j := r; j >= i; j-- {
			if s[i] == s[j] {
				i += 1
			} else {
				i = l
				r--
				j = r + 1
			}
		}
		if (r - l + 1) > len(longestStr) {
			longestStr = s[l : r+1]
		}
	}
	return longestStr
}

func main() {
	fmt.Println(longestPalindrome("bacabab"))
	fmt.Println(longestPalindrome("cdghgd"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aa"))
	fmt.Println(longestPalindrome("babcd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("bbbbbbbcd"))
	fmt.Println(longestPalindrome("ccbbbbbbb"))

}
