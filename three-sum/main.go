package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
			}
			if sum < 0 {
				left++
				for nums[left-1] == nums[left] && left < right {
					left++
				}
			} else {
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			}
		}

		for i < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
	threeSum([]int{0, 1, 1})
	threeSum([]int{0, 0, 0})
}
