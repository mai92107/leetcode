package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {

	end := len(s)-1
	start := 0

	for end > start{
		fmt.Println("他是digit嗎",unicode.IsDigit(rune(s[start])))
		fmt.Println("他是letter嗎",unicode.IsLetter(rune(s[start])))
		fmt.Println("他是digit嗎",unicode.IsDigit(rune(s[end])))
		fmt.Println("他是letter嗎",unicode.IsLetter(rune(s[end])))
		if !unicode.IsLetter(rune(s[start])) && !unicode.IsDigit(rune(s[start])){
			start++
			fmt.Printf("%v 不是字母也不是數字 跳過\n",string(s[start]))
			continue
		}
		if !unicode.IsLetter(rune(s[end])) && !unicode.IsDigit(rune(s[end])){
			end--
			fmt.Printf("%v 不是字母也不是數字 跳過\n",string(s[end]))
			continue
		}
		fmt.Printf("第 %v 個 %c 與第 %v 個 %c 相比為 %v \n", start, s[start], end, s[end], strings.EqualFold(string(s[start]),string(s[end])))
		if strings.EqualFold(string(s[start]),string(s[end])){
			start++
			end--
			continue
		}
		return false
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("ab2a"))
}