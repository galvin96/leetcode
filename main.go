package main

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			min := height[j]
			if height[i] <= height[j] {
				min = height[i]
			}
			if min*(j-i) > max {
				max = min * (j - i)
			}
		}
	}
	return max
}

func main() {
	maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	maxArea([]int{1, 1})
}
