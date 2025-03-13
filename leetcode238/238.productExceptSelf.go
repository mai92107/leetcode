package main

import "fmt"

func productExceptSelf(nums []int) []int {
    
	// n(O2)
	// ans := []int{}
	// for i,_ := range nums{
	// 	a := 1
	// 	for j,v := range nums{
	// 		if i == j{
	// 			continue
	// 		}
	// 		a *= v
	// 	}
	// 	ans = append(ans, a)
	// }
	// fmt.Println(ans)
	// return ans

	// n(O)
	//假設nums為[x,y,z,a]
	n := len(nums)
	ans := make([]int,n)

	//左邊走一次
	//ans為[1,x,x*y,x*y*z]
	ans[0] = 1
	for i:=1;i<n;i++{
		ans[i] = nums[i-1]*ans[i-1]
	}


	//右邊走一次
	//ans為[1*y*z*a,x*z*a,x*y*a,x*y*z]
	r := 1
	for i:=n-1;i>=0;i--{
		// ans[i] = 1 * ans[i]  
		ans[i] *= r
		r *= nums[i]
	}
	fmt.Println(ans)
	return ans
}

func main(){
	productExceptSelf([]int{1,2,3,4,5})
}