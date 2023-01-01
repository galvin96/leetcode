package main

func convert(s string, numRows int) string {
	maxDur := (numRows - 1) * 2
	prevDur := maxDur
	nextDur := 0

	row := 0
	res := ""

	if len(s) <= numRows || numRows == 1 {
		return s
	}

	for row < numRows && row < len(s) {
		index := row
		res += string(s[index])
		for index < len(s)-1 {
			index += prevDur
			if prevDur > 0 && index < len(s) {
				res += string(s[index])
			}
			index += nextDur
			if nextDur > 0 && index < len(s) {
				res += string(s[index])
			}
		}

		prevDur -= 2
		nextDur += 2
		row++
	}
	return res
}

func main() {
	convert("ABC", 1)
	convert("AB", 1)
	convert("PAYPALISHIRING", 4)
	convert("PAYPALISHIRING", 3)
	convert("A", 2)
}
