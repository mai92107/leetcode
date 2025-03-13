package main

import (
	"fmt"
	"sort"
)

// // compare one on one
// func longestCommonPrefix(strs []string) string {
//     common := "" // common string
// 	testAl := "" // current testing alphabet
// 	out:	// take first word for example and loop each alphabet to check
// 	for j,_ := range strs[0]{
// 		for i,v := range strs{
// 			// leave if current index is out of length
// 			if j >= len(v){
// 				break out
// 			}
// 			// set standard
// 			if i == 0{
// 				testAl = string(v[j])
// 				continue
// 			}
// 			// if same as standard, then compare next
// 			if testAl == string(v[j]){
// 				continue
// 			}
// 			// if not the same, then return current answer
// 			if testAl != string(v[j]){
// 				return common
// 			}
// 		}
// 		common += testAl
// 	}
// 	return common
// }

// compare one on one
func longestCommonPrefix(strs []string) string {

	// sort first to ensure the first word is the shortest...
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i])<len(strs[j])
	})

	// take first word for example and loop each alphabet to check
	for i,_ := range strs[0]{
		for _,v := range strs{
			// if not equal to the character of first word 
			if strs[0][i] != v[i]{
				// return the answer util current pointer
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
func main(){
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}