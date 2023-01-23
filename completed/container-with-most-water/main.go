package main

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1

	for left < right {
		rHeight := height[right]
		lHeight := height[left]

		min := rHeight
		if lHeight <= rHeight {
			min = lHeight
		}
		if min*(right-left) > max {
			max = min * (right - left)
		}

		if lHeight < rHeight {
			left++
		} else {
			right--
		}
	}
	return max
}

func main() {
	maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	maxArea([]int{1, 1})
}
