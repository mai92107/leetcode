package main


func removeDuplicates(nums []int) int {
    
	k := 0
	kill := false
	
	for i,v := range nums{
		if i == 0 {
			nums[k] = v
			k++
			continue
		}
		if i != 0 && nums[i-1] != v{
			nums[k] = v
			k++
			kill = false
			continue
		}
		if nums[i-1] == v && !kill{
			nums[k] = v
			k++
			kill = true
		}
	}
	return k
}

func main(){
	removeDuplicates([]int{1,1,1,2,2,3})
}