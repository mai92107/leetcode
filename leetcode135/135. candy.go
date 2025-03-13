package main

import "fmt"

func candy(ratings []int) int {

	var candyCount []int = make([]int, len(ratings))
	total := 0 

	//左邊跑一次 當左邊比右邊大就多給1個
	for i:=0;i<len(ratings);i++{
		candyCount[i] = 1

		if i!=0&&ratings[i]>ratings[i-1]{
			candyCount[i] = candyCount[i-1]+1
		}
	}

	//右邊跑一次 當右邊比左邊大就多給1個
	for i:=len(ratings)-1;i>0;i--{
		if ratings[i-1]>ratings[i] && !(candyCount[i-1]>candyCount[i]){
			candyCount[i-1] = candyCount[i]+1
		}
	}
	fmt.Println(candyCount)
	for _,v := range candyCount{
		total += v
	}
	return total
}

func main(){
	fmt.Println(candy([]int{1,2,87,87,87,2,1}))
}

// 1231321