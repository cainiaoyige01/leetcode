package main

/**
 * @Author: Niuzai
 * @Date: 2023/4/24 18:12
 * @Description:两两交换链表中的节点
 */
func main() {

}

// swapPairs 将两两节点视为内部短链接 交互两个节点后返回新的头结点 接上尾巴
func swapPairs(head *ListNode) *ListNode {
	//这里是两种异常的case
	if head == nil && head.Next == nil {
		return head
	}
	newHead := swap(head, head.Next)
	pre := head
	if pre.Next == nil || pre.Next.Next == nil {
		return newHead
	}
	cur := pre.Next
	next := cur.Next
	for cur != nil && next != nil {
		pre.Next = swap(cur, next)
		if cur.Next == nil {
			break
		}
		pre = cur
		cur = cur.Next
		next = cur.Next
	}
	return newHead
}
func swap(cur, next *ListNode) *ListNode {
	cur.Next = next.Next
	next.Next = cur
	return next
}
