package main

import "fmt"

func decimalToBinary(num int) string {
	var binary string

	for num != 0 {
		s := fmt.Sprintf("%d", num%2)
		binary = s + binary
		num = num / 2
	}

	return binary
}

func Solution(N int) int {
	str := decimalToBinary(N)
	max := 0
	i := 0
	for i < len(str)-1 {
		j := i + 1
		for ; j < len(str) && str[j] == '0'; j++ {
		}
		if j-i-1 > max && j < len(str) {
			max = j - i - 1
		}
		i = j
	}
	return max
}

func main() {
	Solution(529)
	Solution(20)
	Solution(32)
	Solution(29)
}
