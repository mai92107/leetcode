package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {

	// 利用 split 確認是否包含
	// result := strings.Split(haystack, needle)
	// if len(result)==1{
	// 	return -1
	// }
	// return len(result[0])

	// 直接用內建func
	return strings.Index(haystack,needle)
}

func main(){
	fmt.Println(strStr("a","a"))
}