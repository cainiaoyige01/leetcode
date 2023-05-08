package main

import "fmt"

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。这是关键点
示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
*/

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))                      // 2
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})) // 5
}

// removeDuplicates 利用双指针完成 也就是快慢指针！注意前提是有序数组了！！！
func removeDuplicates(nums []int) int {
	slow := 0
	fast := 0
	for fast < len(nums)-1 { //注意这里的减一
		if nums[fast] != nums[fast+1] {
			slow++
			fast++
			nums[slow] = nums[fast]
			continue
		}
		fast++
	}
	return slow + 1
}
