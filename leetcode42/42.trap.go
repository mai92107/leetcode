package main

import "fmt"

// func trap(height []int) int {
    
// 	var water []int = make([]int, len(height))
// 	high := 0

// 	// 從左邊開始取最大值曲線
// 	for i,v := range height{
// 		if v > high{
// 			high = v
// 		}
// 		water[i] = high
// 	}
// 	count := 0

// 	// 從右邊開始取最大值曲線
// 	for i:=len(height)-1;i>=0;i--{
// 		if i == len(height)-1{
// 			high = height[i]
// 		}
// 		if height[i] > high{
// 			high = height[i]
// 		}

// 		water[i] = min(high,water[i])
// 	}

// 	// 取左邊加右邊交集(最小值)
// 	for i := 0; i < len(water); i++ {
// 		count += (water[i]-height[i])
// 	}
// 	fmt.Println(water)
// 	return count
// }
// func min(a int, b int)int{
// 	if a < b{
// 		return a
// 	}
// 	return b
// }

// one pass solution
func trap(height []int) int {
    
	left, right := 0,len(height)-1
	left_max,right_max := height[left],height[right]
	water_tank := 0

	for left < right {
		// 找短邊算儲水
		if left_max < right_max{
			water_tank += (left_max-height[left])
			left++
			left_max = max(left_max,height[left])
		}else{
			water_tank += (right_max-height[right])
			right--
			right_max = max(right_max,height[right])
		}
	}
	return water_tank
}
func max(a int, b int)int{
	if a > b{
		return a
	}
	return b
}
func main(){
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}
