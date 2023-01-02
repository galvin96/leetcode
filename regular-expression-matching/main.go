package main

import "fmt"

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}

	if len(p) > 1 && p[1] == '*' {
		if len(s) > 0 && (p[0] == '.' || p[0] == s[0]) {
			return isMatch(s, p[2:]) || isMatch(s[1:], p)
		} else {
			return isMatch(s, p[2:])
		}
	} else {
		if len(s) == 0 {
			return false
		}
		if s[0] == p[0] || p[0] == '.' {
			return isMatch(s[1:], p[1:])
		}

		return false
	}
}

func main() {
	fmt.Println(isMatch("b", "aaa."))          //F
	fmt.Println(isMatch("aaa", "ab*a*c*a"))    //T
	fmt.Println(isMatch("aaa", "a*a"))         // T
	fmt.Println(isMatch("aaa", "a*aa"))        // T
	fmt.Println(isMatch("aaa", "a*aaa"))       // T
	fmt.Println(isMatch("aaadd", "ca*d*"))     // F
	fmt.Println(isMatch("aaadde", "ca*d*"))    // F
	fmt.Println(isMatch("ab", ".*c"))          // F
	fmt.Println(isMatch("aaaaaaa", "a...."))   // F
	fmt.Println(isMatch("aaaaaaa", "a......")) // T
	fmt.Println(isMatch("aa", "a"))            // F
	fmt.Println(isMatch("aa", "a*"))           // T
	fmt.Println(isMatch("aa", "aa"))           // T
	fmt.Println(isMatch("aa", ".."))           // T
	fmt.Println(isMatch("aa", "."))            // F
	fmt.Println(isMatch("aa", ".*"))           // T
}
