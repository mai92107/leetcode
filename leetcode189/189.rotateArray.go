package main

import "fmt"

func rotate(nums []int, k int)  {

	reverse(nums,0,len(nums)-1)
	if k < len(nums){
		reverse(nums,0,k-1)
		reverse(nums,k,len(nums)-1)
	}

	fmt.Println(nums)
}

func reverse(arr []int, s int, e int){
	for s <= e {
		arr[s],arr[e] = arr[e],arr[s]
		s++
		e--
	}
}
func main(){
	rotate([]int{-1},2)
}