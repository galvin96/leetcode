package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	str := strconv.Itoa(x)

	for i, s := range str {
		if int(len(str)/2) == i {
			return true
		}
		if s != int32(str[len(str)-i-1]) {
			return false
		}
	}
	return true
}

func main() {
	isPalindrome(121)
	isPalindrome(1221)
	isPalindrome(123)
	isPalindrome(-121)
}
