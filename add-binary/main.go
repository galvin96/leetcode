package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	carry := 0
	ret := ""
	for a != "" || b != "" || carry == 1 {
		lastA := 0
		if a != "" {
			lastA, _ = strconv.Atoi(string(a[len(a)-1]))
			a = a[:len(a)-1]
		}
		lastB := 0
		if b != "" {
			lastB, _ = strconv.Atoi(string(b[len(b)-1]))
			b = b[:len(b)-1]
		}

		sum := lastA + lastB
		if sum == 2 {
			if carry == 1 {
				sum = 1
			} else {
				sum = 0
			}
			carry = 1
		} else if sum == 1 {
			if carry == 1 {
				sum = 0
			}
		} else {
			sum += carry
			carry = 0
		}
		ret = fmt.Sprint(sum) + ret
	}
	return ret
}

func main() {
	addBinary("11", "11")
	addBinary("11", "1")
	addBinary("1010", "1011")
}
