package main

import (
	"sort"
)

func hIndex(citations []int) int {
    

	sort.Slice(citations,func(i int, j int) bool{
		return citations[i]>citations[j]
	})
	for i,v := range citations{
		if i >= v{
			return i
		}
	}
	return len(citations)
}
func main(){
	hIndex([]int{100})
}