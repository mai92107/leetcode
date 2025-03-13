package main

import (
	"fmt"

	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	// nums1 = nums1[:m]
    // nums1 = append(nums1, nums2...)
	// sort.Slice(nums1, func(i, j int) bool {
	// 	return nums1[i]<nums1[j]
	// })
	k := m
	for i:=0;i<n;i++{
		nums1[k] = nums2[i]
		k++
	}
	sort.Ints(nums1)
	fmt.Println(nums1)
}
func main(){
	merge([]int{1,2,3,0,0,0},3,[]int{2,5,6},3)
}
