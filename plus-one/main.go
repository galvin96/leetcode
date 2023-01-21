package main

func plusOne(digits []int) []int {
	i := len(digits) - 1

	for ; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}

	if i == -1 {
		return append([]int{1}, digits...)
	}

	return digits
}

func main() {
	plusOne([]int{9})
	plusOne([]int{9, 9})
	plusOne([]int{1, 2, 3, 4})
	plusOne([]int{6, 7, 8, 9})
}
