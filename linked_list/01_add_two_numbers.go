package main

/**
 * @Author: Niuzai
 * @Date: 2023/4/24 8:47
 * @Description:两数相加
 */
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	//首先进行遍历两个链表
	for l1 != nil || l2 != nil {
		//建立两个临时存值的变量
		n1, n2 := 0, 0
		//分别取出l1、l2中的值
		if l1 != nil {
			//取值
			n1 = l1.Val
			//取到值后往后走一格
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//统计总值sum carry 是上一次溢出的值
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		//加入到新的链表中去
		if head == nil {
			//创建一个head 并为其赋值
			head = &ListNode{Val: sum}
			//head=head.Next 如果不适用一个额外的tail来接受的哇 那么head又变成nil了
			tail = head
		} else {
			//为第一次没有tail没有往下指 下一格赋值
			tail.Next = &ListNode{Val: sum}
			//指针走到下一格
			tail = tail.Next
		}
	}
	//如果到了这里carry还是大于0的哇 就往后在添加一个链格
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}

	}

	return head
}
