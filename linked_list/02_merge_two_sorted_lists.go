package main

/**
 * @Author: Niuzai
 * @Date: 2023/4/24 13:23
 * @Description:合并两个有序链表
 */
func main() {

}

// mergeTwoLists 利用中间变量进行转递
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	temp1 := &ListNode{}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp1.Val = l1.Val
			temp1.Next = l2
			l1 = l1.Next
		} else {
			temp1.Val = l2.Val
			temp1.Next = l1
			l2 = l2.Next
		}
		temp1 = temp1.Next
	}
	if l1 != nil {
		temp1.Val = l1.Val
		temp1.Next = l1.Next
	}
	if l2 != nil {
		temp1.Val = l2.Val
		temp1.Next = l2.Next
	}
	return temp1
}
