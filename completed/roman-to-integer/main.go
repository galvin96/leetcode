package main

import "strings"

func romanToInt(s string) int {
	num := 0

	conversions := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romanNumeral := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	i := len(romanNumeral) - 1

	for len(s) > 0 {
		if strings.HasPrefix(s, romanNumeral[i]) {
			s = s[len(romanNumeral[i]):]
			num += conversions[i]
		} else {
			i--
		}
	}

	return num
}

func main() {
	romanToInt("III")
	romanToInt("LVIII")
	romanToInt("MCMXCIV")
}
