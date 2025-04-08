package main

import (
	"fmt"
)
// func twoSum(numbers []int, target int) []int {
    
// 	var ans []int = make([]int, 2)

// 	for i,v := range numbers{
// 		if contains(numbers,i+1,target-v){
// 			ans[0] = i+1
// 			idx := index(numbers,i+1,target-v)
// 			ans[1] = idx
// 			break
// 		}
// 	}
// 	return ans
// }

// func contains(arr []int, start, target int)bool{
// 	for i:=start;i<len(arr);i++{
// 		if arr[i] == target{
// 			return true
// 		}
// 	}
// 	return false
// }
// func index(arr []int,start, target int)int{
// 	for i:=start;i<len(arr);i++{
// 		if arr[i] == target{
// 			return i+1
// 		}
// 	}
// 	return 0
// }

func twoSum(numbers []int, target int) []int {

	l := 0
	r := len(numbers)-1

	for l<r{
		sum := numbers[l]+numbers[r]
		if sum>target{
			r--
		}
		if sum<target{
			l++
		}
		if sum==target{
			break
		}
	}
	return []int{l+1,r+1}
}
func main(){
	fmt.Println(twoSum([]int{0,0,3,4},0))
}