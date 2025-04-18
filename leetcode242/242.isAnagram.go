package main

import (
	"fmt"
)
func isAnagram(s string, t string) bool {

	if len(s) != len(t){
		return false
	}
    
	com := map[rune]int{}

	for _,v := range s{
		com[v]++
	}

	fmt.Println(com)

	for _,v := range t{
		count,exist:=com[v]
		if !exist ||  count == 0{
			return false
		}
		com[v]--
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram","nagaram"))
}