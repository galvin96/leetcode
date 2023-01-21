package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			ret++
		}
	}
	return ret
}

func main() {
	removeDuplicates([]int{1, 1, 2})
	removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}
