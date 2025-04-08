package main

import "fmt"

func maxArea(height []int) int {
    
	l := 0
	r := len(height)-1
	max := 0

	for l<r{
		area := min(height[l], height[r]) * (r - l)

		// 紀錄大的
		if area > max{
			max = area
		}

		// 比較短的往內移動
		if height[l]<height[r]{
			l++
		}else{
			r--
		}
	}
	return max
}

func min(a,b int)int{
	if a<b{
		return a
	}else{
		return b
	}
}

func main(){
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}