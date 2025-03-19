package main

import (
	"fmt"
)
func isValid(s string) bool {

	var stack []string
	compare := map[string]string{
		")":"(",
		"}":"{",
		"]":"[",
	}
	for _,v := range s{
		if string(v) == "(" || string(v) == "[" || string(v) == "{"{
			stack = append(stack,string(v))
			continue
		}
		if len(stack)!=0 && stack[len(stack)-1] == compare[string(v)]{
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	if len(stack) == 0{
		return true
	}
	return false
}


func main(){
	fmt.Println(isValid("()[]{}"))
}