package main

import "strconv"

func reverse(x int) int {
	res := ""

	if x < 0 {
		x = -x
		res = "-"
	}

	for x > 0 {
		res += strconv.Itoa(x % 10)
		x /= 10
	}

	num, err := strconv.ParseInt(res, 10, 32)
	if err != nil {
		return 0
	}

	return int(num)
}

func main() {
	reverse(1534236469)
	reverse(123)
	reverse(-123)
	reverse(120)

}
