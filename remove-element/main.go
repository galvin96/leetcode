package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)
	removeElement([]int{3, 2, 2, 3}, 2)
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
}
