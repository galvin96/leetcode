package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := int(math.Abs(float64(nums[0] + nums[1] + nums[2] - target)))
	ret := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if int(math.Abs(float64(sum-target))) < min {
				min = int(math.Abs(float64(sum - target)))
				ret = sum
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
	}
	return ret
}

func main() {
	threeSumClosest([]int{0, 1, 2}, 3)
	threeSumClosest([]int{-1, 2, 1, -4}, 1)
	threeSumClosest([]int{0, 0, 0}, 1)
}
