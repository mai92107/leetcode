package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int)(ans []string) {
	ans = []string{""}
	index := 0
	for i,v := range words{
		ans[index]+=(v+" ")

		// 非最後一個 且 加上下一個字會爆 換行前進入
		if i != len(words)-1 && len(ans[index])+len(words[i+1])>maxWidth{
			totalSpace,newWordArr := parseData(&ans[index],maxWidth)
			//若有空格需分配則
			//  重新調整當前句子
			if totalSpace != 0{
				ans[index] = getFullStringWithPadding(newWordArr,totalSpace,len(newWordArr)-1)
			}
			//換行
			ans = append(ans, "")
			index++
		}
		//處理最後一句話 不重新分配空格 只補空格在最後
		if i == len(words)-1{
			totalSpace,newWordArr := parseData(&ans[index],maxWidth)
			ans[index] += getStringSpace(totalSpace-len(newWordArr)+1)
		}
	} 

	for _,v := range ans{
		fmt.Println(v)
	}
	return
}

func parseData(origin *string, maxWidth int)(space int,arr []string){
	*origin = (*origin)[:len(*origin)-1]
	space = maxWidth - len(strings.TrimSpace(*origin))
	arr = strings.Split(strings.TrimSpace(*origin)," ")
	space += (len(arr)-1)
	return
}

func getFullStringWithPadding(arr []string, totalSpace int, gapCount int)(ans string){
	ans += arr[0]
	if gapCount == 0{
		ans += getStringSpace(totalSpace)
	}
	for i:=0;i<gapCount;i++{
		ans += getStringSpace(countSpace(totalSpace,gapCount,i))
		ans += arr[i+1]
	}
	return
}

func countSpace(totalSpace int, gapCount int, no int)(ans int){
	ans = totalSpace / gapCount
	if totalSpace % gapCount > no{
		ans += 1
	}
	return
}

func getStringSpace(count int)(ans string){
	
	for i:=0;i<count;i++{
		ans+=" "
	}
	return
}
func main(){
	fmt.Println(fullJustify([]string{"What","must","be","acknowledgment","shall","be"},16))
}
