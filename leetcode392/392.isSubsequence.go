package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {

	si := 0
	ti := 0

	for si!=len(s)&&ti!=len(t){
		if s[si]==t[ti]{
			si++
		}
		ti++
	}
	if si!=len(s){
		return false
	}
	return true
}

func main() {
	fmt.Println(isSubsequence("axc","ahbgdc"))
}