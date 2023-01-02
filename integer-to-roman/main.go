package main

func intToRoman(num int) string {
	str := ""

	conversions := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romanNumeral := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	i := len(conversions) - 1

	for num > 0 {
		if num >= conversions[i] {
			num -= conversions[i]
			str += romanNumeral[i]
		} else {
			i--
		}
	}

	return str
}

func main() {
	intToRoman(3)
	intToRoman(58)
	intToRoman(1994)
	intToRoman(3999)
}
