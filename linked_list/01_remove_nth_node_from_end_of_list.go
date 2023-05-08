package main

/**
 * @Author: Niuzai
 * @Date: 2023/4/24 13:01
 * @Description:删除链表的倒数第N个结点
 */
func main() {

}

// removeNthFromEnd 这个方法行不了  关键点是无法删除头节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var tail *ListNode
	var tail1 *ListNode
	tail = head
	tail1 = head
	count := 0
	for tail != nil {
		count++
		tail = tail.Next
	}
	if count == 1 {
		return tail
	}
	temp := count - n
	//n1:=0
	for i := 0; i < temp-1; i++ {
		tail1 = tail1.Next
	}
	tail1.Next = tail1.Next.Next
	return head

}

// removeNthFromEnd 关键点是既要处理头结点 有要防止头结点丢失  要借助哑节点
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Val: 0, Next: head}
	front, rear := dummyNode, dummyNode
	for i := 0; i <= n; i++ {
		rear = rear.Next
	}
	//这是一个很妙的做法 利用已知点来求出差值
	for rear != nil {
		rear = rear.Next
		front = front.Next
	}
	front.Next = front.Next.Next
	return dummyNode.Next
}
