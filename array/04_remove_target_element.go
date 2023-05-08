package main

import (
	"fmt"
)

/*
*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。
例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
*/

// 题意是 不能使用额外的数组
func main() {
	fmt.Println(removeElement1([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement1([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

// removeElement 这是利用双指针来解决的
func removeElement(nums []int, val int) int {
	n := len(nums)
	left := 0
	right := n - 1
	for left < right {
		switch {
		case nums[right] == val:
			right--
		case nums[left] != val:
			left++
		default:
			nums[left] = nums[right]
			nums[right] = val
			left++
			right--
		}
	}
	return left + 1
}

// removeElement1 这是官方简单的解决 利用的快慢指针的原理
func removeElement1(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		//同一起跑也没事
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
