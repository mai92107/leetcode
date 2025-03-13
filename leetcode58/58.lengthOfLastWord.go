package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
	c := strings.Split(s, " ")
	lc := strings.Split(c[len(c)-1], "")
	return len(lc)
}

func main(){
	fmt.Println(lengthOfLastWord("Hello World"))
}
