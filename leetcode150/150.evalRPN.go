package main

import (
	"fmt"
	"strconv"
	// "strconv"
)

func evalRPN(tokens []string) int {
	numStack := []int{}
	for _,v := range tokens{

		// 使用map對應取代switch
		operation := map[string]func(a,b int)int{
			"+":func(a,b int) int {return a+b},
			"-":func(a,b int) int {return a-b},
			"*":func(a,b int) int {return a*b},
			"/":func(a,b int) int {return a/b},
		}
		if cal,exist := operation[v];exist{
			a,b := numStack[len(numStack)-2],numStack[len(numStack)-1]
			numStack = append(numStack[:len(numStack)-2], cal(a,b))
		}else{
			num,_ := strconv.Atoi(v)
			numStack = append(numStack, num)
		}

		// //使用switch對應＋-*/的算法
		// // 若為符號則從棧中取出數字來運算，算完再壓入棧中模擬括弧中運算結果
		// switch v{
		// case "+":
		// 	a,b := numStack[len(numStack)-2],numStack[len(numStack)-1]
		// 	numStack = numStack[:len(numStack)-2]
		// 	numStack = append(numStack, a+b)
		// 	// fmt.Printf("%v + %v = %v \n",a,b,a+b)
		// case "-":
		// 	a,b := numStack[len(numStack)-2],numStack[len(numStack)-1]
		// 	numStack = numStack[:len(numStack)-2]
		// 	numStack = append(numStack, a-b)
		// 	// fmt.Printf("%v - %v = %v \n",a,b,a-b)
		// case "*":
		// 	a,b := numStack[len(numStack)-2],numStack[len(numStack)-1]
		// 	numStack = numStack[:len(numStack)-2]
		// 	numStack = append(numStack, a*b)
		// 	// fmt.Printf("%v * %v = %v \n",a,b,a*b)
		// case "/":
		// 	a,b := numStack[len(numStack)-2],numStack[len(numStack)-1]
		// 	numStack = numStack[:len(numStack)-2]
		// 	numStack = append(numStack, a/b)
		// 	// fmt.Printf("%v / %v = %v \n",a,b,a/b)

		// // 數字的情況 
		// // 只要是數字就往棧中壓入
		// default:
		// 	num,_ := strconv.Atoi(v)
		// 	numStack = append(numStack, num)
		// }
	}
	return numStack[0]
}



func main(){
	fmt.Println(evalRPN([]string{"4","13","5","/","+"}))
}