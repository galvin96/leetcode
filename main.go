package main

import "sort"

func threeSum(nums []int) [][]int {
	var ret [][]int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					set := []int{nums[i], nums[j], nums[k]}
					sort.Ints(set)
					ret = append(ret, set)
				}
			}
		}
	}
	return ret
}

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
	threeSum([]int{0, 1, 1})
	threeSum([]int{0, 0, 0})
}
