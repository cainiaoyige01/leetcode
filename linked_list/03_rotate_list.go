package main

/**
 * @Author: Niuzai
 * @Date: 2023/4/25 9:54
 * @Description:旋转链表
 */
func main() {

}

// rotateRight 首先要理解 head就是一个地址而已 它能形成链表是因为 *ListNode 的指针指向下一个地址 让你能找到而已
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := Count(head)
	k = k % n
	if k == 0 {
		return head
	}
	rear := head
	fast := head
	for i := 0; i < k; i++ {
		rear = rear.Next
	}
	for rear.Next != nil {
		fast = fast.Next
		rear = rear.Next
	}
	sum := fast.Next
	fast.Next = nil
	cur := sum
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
	return sum
}
func Count(head *ListNode) int {
	temp := head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	return count
}
