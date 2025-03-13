package main

import "fmt"

func jump(nums []int) int {

	count := 1
	curEnd := 0
	farthest := 0

	for i:=0;i<len(nums)-1;i++{

		//保存當前步內接下來最遠可以到哪裡
		farthest = max(farthest,i+nums[i])
		
		//到達當前步邊界
		if i == curEnd{
			//步數加一
			count++
			//更新當前步可到達邊界
			curEnd = farthest

			//到達終點
			if curEnd >= len(nums)-1{
				break
			}
		}
	}

	fmt.Println(count)
	return count
}
func max(a int, b int)int{
	if a > b{
		return a
	} 
	return b
}

func main(){
	jump([]int{2,3,1,1,4})
}
