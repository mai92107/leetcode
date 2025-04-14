package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	// 沒資料就回 0
	if len(nums) == 0{
		return 0
	}
    // 先排序
	sort.Ints(nums)
	// 只要有資料 最少長度就會是 1
	count := 1
	maxCount := 1

	for i,v := range nums{
		if i == 0{
			continue
		}
		// 正確排序
		if v - nums[i-1]==1{
			count++
		// 重複跳過
		}else if v == nums[i-1]{
			continue
		// 斷掉了則比較目前連續數有沒有比紀錄的最大連續數多
		}else{
			if maxCount < count{
				maxCount = count
			}
			// 初始化
			count = 1
		}
	}
	if maxCount < count{
		maxCount = count
	}
	return maxCount
}

func main() {
	fmt.Println(longestConsecutive([]int{100,4,200,1,3,2}))
}