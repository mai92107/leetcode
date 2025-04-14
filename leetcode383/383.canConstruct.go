package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	
	count := map[rune]int{}
    
	for _,v := range magazine{
		count[v]++
	}

	for _,v := range ransomNote{
		if count[v] == 0{
			return false
		}
		count[v]--
	}

	return true
}

func main(){
	fmt.Println(canConstruct("aa","aab"))
}