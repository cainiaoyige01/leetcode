package main

import "fmt"

/*
*
  - @Author: Niuzai
  - @Date: 2023/4/9 11:20
  - @Description:下一个排序
    实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
    如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
    必须 原地 修改，只允许使用额外常数空间。
    示例 1：
    输入：nums = [1,2,3]
    输出：[1,3,2]
*/
func main() {
	nums := []int{1, 2, 7, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums) // [1 3 1 2 4 7] // bingo
}

// nextPermutation 利用双循环进行查找相邻的数 点睛之笔在于数组的反转
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

// reverse 数组进行反转
func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
