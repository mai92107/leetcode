package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {

	arr := strings.Split(s," ")
	if len(arr)!=len(pattern){
		return false
	}
    
	comA := map[rune]string{}
	comB := map[string]rune{}

	for i,v := range pattern{
		valueA,existA := comA[v]
		valueB,existB := comB[arr[i]]

		if existA && valueA!=arr[i]{
			return false
		}
		comA[v] = arr[i]

		if existB && valueB!=v{
			return false
		}
		comB[arr[i]] = v
	}
	return true
}

func main(){
	fmt.Println(wordPattern("abba","dog cat cat dog"))
}