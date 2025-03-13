package main

import (
	"fmt"
)

func intToRoman(num int) string {
	ans := ""
	in := []int{1,4,5,9,10,40,50,90,100,400,500,900,1000}
	s := []string{"I","IV","V","IX","X","XL","L","XC","C","CD","D","CM","M"}
	
	for i := len(in)-1 ; i >= 0 ; i-- {
		for num >= in[i]{
			fmt.Println(s[i],in[i],num)
			ans += s[i]
			num -= in[i]
		}
	}
	return ans
}

func main(){
	fmt.Println(intToRoman(3749))
}