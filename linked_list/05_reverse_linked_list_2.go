package main

/**
 * @Author: Niuzai
 * @Date: 2023/4/25 11:14
 * @Description:反转链表2
 */
func main() {

}

func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	//这里也很妙 到那个点的前一个 而那个点就额外来接受 就可以排除当m=1的时候了也就是头结点
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	rightNode := pre
	for i := 0; i < n-m+1; i++ {
		rightNode = rightNode.Next
	}
	leftPre := pre.Next
	curr := rightNode.Next
	pre.Next = nil
	rightNode.Next = nil
	prepend(leftPre)
	//这里很妙呀 去最后一个地址 她们都是关联的 并最后一个反转就可以得到了反转的链表所在
	pre.Next = rightNode
	leftPre.Next = curr
	return dummy.Next

}

// prepend 使用头插法完成链表的反转
func prepend(head *ListNode) {
	////借助哑节点
	//dummy := &ListNode{Next: head}
	//for head != nil {
	//	cur := dummy.Next
	//	dummy.Next = head
	//	dummy.Next.Next = cur
	//	head = head.Next
	//}
	//return dummy.Next
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
}
