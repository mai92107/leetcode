package main

import (
	"fmt"
)

func isIsomorphic(s, t string)bool{

	arr1 := make([]int,256)
	arr2 := make([]int,256)

	for i,_ := range s{
		if arr1[s[i]] != arr2[t[i]]{
			return false
		}
		arr1[s[i]] = i+1
		arr2[t[i]] = i+1
	}

	return false
}

// func isIsomorphic(s string, t string) bool {
    
// 	if len(s)!=len(t){
// 		return false
// 	}
// 	// 檢查 a 向 對應是否達成
// 	compareMap := map[rune]rune{}
// 	for i:=0;i<len(s);i++{
// 		value,exist:=compareMap[rune(s[i])]
// 		if !exist{
// 			compareMap[rune(s[i])] = rune(t[i])
// 			continue
// 		}
// 		if exist && value == rune(t[i]) {
// 			continue
// 		}
// 		return false
// 	}
// 	// double check 檢查 b 向對應
// 	compareMap = map[rune]rune{}
// 	s,t = t,s
// 	for i:=0;i<len(s);i++{
// 		value,exist:=compareMap[rune(s[i])]
// 		if !exist{
// 			compareMap[rune(s[i])] = rune(t[i])
// 			continue
// 		}
// 		if exist && value == rune(t[i]) {
// 			continue
// 		}
// 		return false
// 	}
// 	return true
// }

func main(){
	fmt.Println(isIsomorphic("badc","baba"))
}