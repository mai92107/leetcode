package main

import (
)



type ListNode struct {
    Val int
    Next *ListNode
}
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 紀錄 鏈 的最原點
	head := &ListNode{}
	curr := head
	ten := 0

	for l1!=nil||l2!=nil||ten!=0{

		// 紀錄l1 l2 的 val, 當為 nil 時才會有初始值 0
		var v1,v2 int

		if l1!=nil{
			v1=l1.Val
			l1=l1.Next
		}
		if l2!=nil{
			v2=l2.Val
			l2=l2.Next
		}

		sum := v1+v2+ten
		ten = sum/10
		digit := sum%10
	
		// 將 node 結構 放入 next
		curr.Next = &ListNode{Val: digit}
		// 將 next 的節點 與目前的鏈接上
		curr = curr.Next	

		// 示意圖
		// 其中 curr.Next = &ListNode{Val: digit} 即為 產生結構[6]
		//	curr = curr.Next 接上鏈

		// head -> [2] -> [4] -> nil
		// head -> [2] -> [4] -> [6] -> nil
	}
	return head.Next
}