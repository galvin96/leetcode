package main

func strStr(haystack string, needle string) int {
	for i := range haystack {
		if i <= len(haystack)-len(needle) && haystack[i] == needle[0] {
			j := 0
			for j < len(needle) && needle[j] == haystack[i+j] {
				j++
			}
			if j == len(needle) {
				return i
			}
		}
	}
	return -1
}

func main() {
	strStr("sadbutsad", "sad")
	strStr("leetcode", "leeto")
}
