package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
    
	pathArr := strings.Split(path,"/")
	pathStack := []string{}
	for _,word := range pathArr{
		if word == "" || word == "."{
			continue
		}
		if word == ".."{
			if len(pathStack) == 0{
				continue
			}
			pathStack = pathStack[:len(pathStack)-1]
		}else{
			pathStack = append(pathStack,word)
		}
	}
	return "/"+strings.Join(pathStack,"/")
}

func main(){
	fmt.Println(simplifyPath("/.../a/../b/c/../d/./"))
}
