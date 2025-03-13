package main

func removeDuplicates(nums []int) int {
	
	k := 0
	t := 999
	for _,v := range nums{

		if v != t{
			nums[k] = v
			k++
		}
		t = v
	}	
	return k
}

func main(){
	removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4})
}