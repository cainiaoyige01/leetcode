package main

import (
	"fmt"
	"sort"
)

/*
*
  - @Author: Niuzai
  - @Date: 2023/4/10 9:11
  - @Description:在排序数组中查找元素的第一个和最后一个位置

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
进阶：
你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
*/
func main() {
	//fmt.Println(searchRange1([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]
	//fmt.Println(searchRange3([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]
	//fmt.Println(searchRange3([]int{1}, 1))                 // [0, 0]
	fmt.Println(searchRange3([]int{}, 0))

}

// searchRange1 尝试使用左右指针来完成
func searchRange1(nums []int, target int) []int {
	n := len(nums)
	left := 0
	right := n - 1
	flag1 := false
	flag2 := false

	numsRange := []int{-1, -1}
	for left <= right {
		if !flag2 {
			if nums[right] == target {
				numsRange[1] = right
				flag2 = true
			} else {
				right--
			}
		}
		if !flag1 {
			if nums[left] == target {
				numsRange[0] = left
				flag1 = true
			} else {
				left++
			}
		}
		if flag1 && flag2 {
			return numsRange
		}

	}
	return numsRange
}

// searchRange2
func searchRange2(nums []int, target int) []int {
	indexLeft := sort.SearchInts(nums, target)
	if indexLeft == len(nums) || nums[indexLeft] != target {
		return []int{-1, -1}
	}
	indexRight := sort.SearchInts(nums, target+1) - 1
	return []int{indexLeft, indexRight}

}

// searchRange3 使用二分法的来解决  关键点是一个一个寻找 先找第一个 再找最后一个
func searchRange3(nums []int, target int) []int {
	//这个关键的哇！
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	first := findFirstNums(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}
	last := findLastNums(nums, target)
	return []int{first, last}
}

// findLastNums 注意在查找多个数中靠那边哪个数时 需要谁来接近
func findLastNums(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
func findFirstNums(nums []int, target int) int {
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
	if target == nums[left] {
		return left
	}
	return -1
}
