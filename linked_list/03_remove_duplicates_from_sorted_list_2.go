package main

/**
 * @Author: Niuzai
 * @Date: 2023/4/25 10:15
 * @Description:删除排序链表中的重复元素
 */
func main() {

}

// deleteDuplicates 注意啦 也可能从头结点就开始删除了 所以要借助哑节点
func deleteDuplicates(head *ListNode) *ListNode {
	//首先判断异常的两种case
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil {
		last := pre.Next
		//找到重复的前一个点
		for last.Next != nil && last.Val == last.Next.Val {
			last = last.Next
		}
		//这才是关键点所在
		if pre.Next != last {
			pre.Next = last.Next
		} else {
			pre = pre.Next
		}
	}
	return dummy.Next
}
