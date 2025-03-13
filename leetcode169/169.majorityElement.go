package main

import (
	"sort"
)


func majorityElement(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]>nums[j]
	})
	return nums[len(nums)/2]
}

func main(){
	majorityElement([]int{3,2,3})
}