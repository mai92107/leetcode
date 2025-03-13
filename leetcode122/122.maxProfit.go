package main

func maxProfit(prices []int) int {
    
	profit := 0
	min := prices[0]
	for i:=1;i<len(prices);i++ {
		if prices[i]<=min{
			min = prices[i]
			continue
		}
		profit += prices[i]-min
		min = prices[i]
	}
	return profit
}

func main(){
	maxProfit([]int{1,2,3,4,5})
}