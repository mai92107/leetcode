package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
    cs := strings.Split(s, "")
	total := 0 
	
	m := map[string]int{
		"I" : 1,
		"V" : 5,
		"X" : 10,
		"L" : 50,
		"C" : 100,
		"D" : 500,
		"M" : 1000,
	}

	for i,v := range cs{
		if i != len(cs)-1 && m[v] < m[cs[i+1]]{
				total -= m[v]
		}else{
			total += m[v]
		}
	}
	return total
}

func main(){
	fmt.Println(romanToInt("MCMXCIV"))
}