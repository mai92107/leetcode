package main

func maxProfit(prices []int) int {
	minI := prices[0]
    max := 0
	for _,v := range prices{
		if minI <= v{
			if v-minI > max {
				max = v-minI
			}
			continue
		}
		minI = v
	}
	return max
}


func main(){
	maxProfit([]int{7,1,5,3,6,4})
}