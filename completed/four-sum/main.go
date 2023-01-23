package main

import (
	"sort"
)

func threeSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
			}
			if sum < target {
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

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-3; i++ {
		newTarget := target - nums[i]

		threeSumArr := threeSum(nums[i+1:], newTarget)
		if threeSumArr != nil {
			for _, arr := range threeSumArr {
				ret = append(ret, append([]int{nums[i]}, arr...))
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}

func main() {
	fourSum([]int{-1, -5, -5, -3, 2, 5, 0, 4}, -7) // [[-5, -5, -1, 4],[-5, -3, -1, 2]]
	fourSum([]int{5, 5, 3, 5, 1, -5, 1, -2}, 4)    // [[-5, 1, 3, 5]]
	fourSum([]int{-3, -1, 0, 2, 4, 5}, 2)          // [[-3, -1, 2, 4]]
	fourSum([]int{0, 0, 0, 0}, 1)                  // []
	fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0)      // [[-2,-1,1,2],[-1,-1,1,1]]
	fourSum([]int{1, 0, -1, 0, -2, 2}, 0)          // [[-2,-1,1,2],[-2, 0, 0, 2],[-1,0,0,1]]
	fourSum([]int{2, 2, 2, 2, 2}, 8)               // [[2,2,2,2]]
}
