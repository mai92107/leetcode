package main


 type ListNode struct {
    Val int
	Next *ListNode
 }

 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	for list1!=nil||list2!=nil{
	
		// 比大小 ， 要的值 插入並取下一個
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next // 讓 cur 移動到新加入的節點
	}
	
	// 如果還有剩餘的節點，直接接到 cur.Next
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	
	return head.Next
 }