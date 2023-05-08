package main

import "fmt"

/*
*
  - @Author: Niuzai
  - @Date: 2023/4/9 11:30
  - @Description:搜索旋转排序数组  使用

整数数组 nums 按升序排列，数组中的值 互不相同 (也就是说每一个数都是唯一的)。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ...,
nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
要求时间复杂度为Log(N)
示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
*/
func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
}

// search 使用类二分法来寻找一个数
func search(nums []int, target int) int {
	//二分法那就是利用左右的是否有序的条件 记住二分法是基于有序的数组
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[left] <= nums[mid]: //假设左边是有序的
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		case nums[mid] <= nums[right]: //设右边是有序的
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}
	}
	return -1
}
