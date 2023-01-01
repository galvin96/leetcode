package main

import "math"

func myAtoi(s string) int {
	i := 0
	num := 0

	for i < len(s) && s[i] == ' ' {
		i++
	}

	if i == len(s) {
		return 0
	}

	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		num = num*10 + int(s[i]-'0')
		i++
		if num > math.MaxInt32 {
			break
		}
	}

	if num > math.MaxInt32 {
		if sign == -1 {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}

	return num * sign
}

func main() {
	myAtoi("9223372036854775808")
	myAtoi("   ")
	myAtoi("")
	myAtoi("-91283472332")
	myAtoi("4193 with words")
	myAtoi("   4200")
	myAtoi("   -4200")
	myAtoi("   000")
	myAtoi("   0002")
}
