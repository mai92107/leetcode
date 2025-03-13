package main

import "math/rand"

type RandomizedSet struct {
    dataMap map[int]int
	dataArr []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
		dataMap: make(map[int]int),
		dataArr: make([]int, 0),
	}
}


func (this *RandomizedSet) Insert(val int) bool {
    
	if _,exist := this.dataMap[val];exist{
		return false
	}

	this.dataArr = append(this.dataArr, val)
	this.dataMap[val] = len(this.dataArr) - 1
	return true	
}


func (this *RandomizedSet) Remove(val int) bool {
	i,exist := this.dataMap[val]
	if !exist{
		return false
	}

	// 把map,arr中最後一個取代要刪的值的位置
	last := this.dataArr[len(this.dataArr)-1]
	this.dataArr[i] = last
	this.dataMap[last] = i

	// 刪arr最後一個
	this.dataArr = this.dataArr[:len(this.dataArr)-1]
	// 刪map最後一個
	delete(this.dataMap,val)

	return true	
}


func (this *RandomizedSet) GetRandom() int {
    
	return this.dataArr[rand.Intn(len(this.dataArr))]
}
