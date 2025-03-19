package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	arr := strings.Fields(s)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
		
	return strings.Join(arr," ")
}

func main(){
	fmt.Println(reverseWords("a good   example"))
}