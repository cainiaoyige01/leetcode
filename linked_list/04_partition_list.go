package main

/**
 * @Author: Niuzai
 * @Date: 2023/4/25 10:42
 * @Description:分隔链表
 */
func main() {

}

// partition 创建额外的组合 注意啦组合的赋值
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	dummy1 := &ListNode{Val: 1, Next: head}
	cur := dummy
	last := dummy1
	for head != nil {
		if head.Val < x {
			cur.Next = head
			cur = cur.Next
		} else {
			last.Next = head
			last = last.Next
		}
		head = head.Next
	}
	if last.Next != nil {
		last.Next = nil
	}
	cur.Next = dummy1.Next
	return dummy.Next
}
