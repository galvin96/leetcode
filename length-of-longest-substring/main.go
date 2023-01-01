package longestStringWithoutRepeating

import "strings"

func lengthOfLongestSubstring(s string) int {
	max := 0
	subStr := ""

	for _, char := range s {
		newChar := string(char)
		if strings.Contains(subStr, newChar) {
			subStr = subStr[strings.Index(subStr, newChar)+1:] + newChar
		} else {
			subStr += newChar
		}

		if len(subStr) > max {
			max = len(subStr)
		}
	}
	return max
}
