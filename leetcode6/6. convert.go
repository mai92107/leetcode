package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1{
		return s
	}
	ans := make([]string,numRows)
	status := false
	i := 0
	k:=0
	for i<len(s){
		if !status{
			ans[k] += string(s[i])
			k++
		}else{
			ans[k] += string(s[i])
			k--
		}
		i++
		if k == 0 || k == numRows-1{
			status = !status
		}
	}
	return strings.Join(ans,"")
}

func main(){
	fmt.Println(convert("PAYPALISHIRING",3))
}