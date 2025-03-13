package main

func canJump(nums []int) bool {

	maxIndex := nums[0]
	for i,v := range nums{
		if maxIndex < i{
			return false
		}
		if maxIndex >= len(nums){
			return true
		}
		maxIndex = Max(maxIndex,i+v)
	}
	return true
}

func Max(a int, b int)int{
	if a > b{
		return a
	} 
	return b
}
func main(){
	canJump([]int{3,2,1,0,4})
}
