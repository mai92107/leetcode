package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	ans := [][]int{}

	// 固定中間值 i，透過 l 與 R 來找
	for i:=0;i<len(nums);i++{
		l,r := i+1,len(nums)-1

		// 固定值 若重複則跳過
		if i>0 && nums[i] == nums[i-1]{
			continue
		}

		for l<r{
			sum := nums[i]+nums[l]+nums[r]

			switch{
			case sum == 0:
				// 符合條件 加入陣列
				ans = append(ans, []int{nums[i],nums[l],nums[r]})
				
				// 若相同跳過
				for l<r && nums[l]==nums[l+1]{
					l++
				}
				for l<r && nums[r]==nums[r-1]{
					r--
				}
				l++
				r--
			case sum > 0:
				r--
			case sum < 0:
				l++
			}
		}
	}
	return ans
}

func main(){
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}