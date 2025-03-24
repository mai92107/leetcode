package main

import (
	"fmt"
	"math"
)

// 建構子
type MinStack struct {
    Stack []int
	Min	  []int
}


func Constructor() MinStack {
    return MinStack{
		Stack: make([]int, 0),
		Min	 : []int{math.MaxInt64},
	}
}


func (this *MinStack) Push(val int)  {
    this.Stack = append(this.Stack, val)
	if this.Min[len(this.Min)-1] >= val{
		this.Min = append(this.Min, val)
	}
}


func (this *MinStack) Pop()  {
	if this.Min[len(this.Min)-1] == this.Stack[len(this.Stack)-1]{
		this.Min = this.Min[:len(this.Min)-1]
	}
    this.Stack = this.Stack[:len(this.Stack)-1]
}


func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main(){
	obj := Constructor();
	obj.Push(5);
	obj.Push(-5);
	obj.Push(6);
	obj.Push(20);
	obj.Push(4);
	obj.Pop();
	param_3 := obj.Top();
	param_4 := obj.GetMin();
	fmt.Println(param_3,param_4)
}