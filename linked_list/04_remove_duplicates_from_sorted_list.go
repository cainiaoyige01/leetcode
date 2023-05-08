package main

/**
 * @Author: Niuzai
 * @Date: 2023/4/25 10:29
 * @Description:删除排序链表中的重复元素
 */
func main() {

}

// deleteDuplicates1 利用了当前节点衍生出来的下一一节去跑 然后再并接
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	for pre != nil {
		last := pre.Next
		for last != nil && last.Val == pre.Val {
			last = last.Next
		}
		pre.Next = last
		pre = pre.Next
	}
	return head
}
