package main

import "fmt"

/*
*
  - @Author: Niuzai
  - @Date: 2023/4/10 11:05
  - @Description:搜索插入位置

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。
示例 1:
输入: [1,3,5,6], 5
输出: 2
*/
func main() {
	fmt.Println(searchInsert1([]int{1, 3, 5, 6}, 5)) //2
	fmt.Println(searchInsert1([]int{1, 3, 5, 6}, 2)) //1
	fmt.Println(searchInsert1([]int{1, 3, 5, 6}, 0)) //0
}

// searchInsert 使用二分法来解决
func searchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}

	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid
		} else {
			right = mid - 1
		}
	}

	if nums[left] >= target {
		return left
	}
	return left + 1
}

// searchInsert1 优化后的二分法
func searchInsert1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		default:
			return mid
		}
	}
	return left
}
