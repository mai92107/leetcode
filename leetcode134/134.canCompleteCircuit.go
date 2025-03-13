package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	
	//暴力解法
	// //起始點位置
	// for i:=0;i<len(gas);i++{
	// 	count := 1
	// 	//開始計算
	// 	j := i
	// 	gaso := gas[j] - cost[j]
	// 	for gaso >= 0 && count < len(gas){
	// 		count ++
	// 		j++
	// 		j = j%len(gas)
	// 		gaso += gas[j]
	// 		gaso -= cost[j]
	// 	}
	// 	if gaso >= 0 && count == len(gas){
	// 		return i
	// 	}
	// }
	// return -1

	// greedy solution:
	cur_sum := 0
	total_sum := 0
	start := 0
	for i:=0;i<len(gas);i++{
		sum := gas[i]-cost[i]
		cur_sum += sum
		total_sum += sum
		if cur_sum < 0{
			cur_sum = 0
			start = i+1
		}
	}
	if total_sum < 0{
		return -1
	}
	return start

}

func main(){
	fmt.Println(canCompleteCircuit([]int{1,2,3,4,5,5,70},[]int{2,3,4,3,9,6,2}))
}