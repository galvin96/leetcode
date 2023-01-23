package main

import "math"

func divide(dividend int, divisor int) int {
	ret := dividend / divisor
	if ret >= int(math.Pow(2, 31)) {
		return int(math.Pow(2, 31)) - 1
	}
	return ret
}

func main() {
	divide(10, 3)
	divide(7, -3)
}
