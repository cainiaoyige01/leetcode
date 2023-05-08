package main

/**
 * @Author: Niuzai
 * @Date: 2023/4/25 12:23
 * @Description:环形链表
 */
func main() {

}

// hasCycle 可以使用快慢指针来做 也可以是用哈希来做 存地址
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	fast = fast.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
