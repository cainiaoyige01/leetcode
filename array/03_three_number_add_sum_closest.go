package main

import (
	"fmt"
	"math"
	"sort"
)

/*
*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
示例：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
*/
func main() {
	fmt.Println(threeSumClosest1([]int{-1, 2, 1, -4}, 1))                 // 2
	fmt.Println(threeSumClosest1([]int{0, 2, 1, -3}, 1))                  // 0
	fmt.Println(threeSumClosest1([]int{1, 2, 4, 8, 16, 32, 64, 128}, 84)) // 84
}

// threeSumClosest 利用自己的想法来做 增加额外的空间
func threeSumClosest(nums []int, target int) int {
	//想对数组进行排序 有利于操作
	sort.Ints(nums)
	var nums2 []int
	map1 := make(map[int]int, len(nums))
	for index, value := range nums {
		left := index + 1
		right := len(nums) - 1
		for left < right {
			temp := value + nums[left] + nums[right]
			//fmt.Println(temp)
			abs := int(math.Abs(float64(target-temp))) + 1
			nums2 = append(nums2, abs)
			map1[abs] = temp
			if temp > target {
				right--
			} else {
				left++
			}
		}

	}
	sort.Ints(nums2)

	return map1[nums2[0]]
}

// threeSumClosest1 官方的解法 关键点是初始化定义了一个超大的数 小优化：去重
func threeSumClosest1(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	minAbs := 1<<31 - 1 //这是关键点 我从没想过这样定义大的数
	minSum := 0
	//fmt.Println(minAbs)
	for index, value := range nums {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		left := index + 1
		right := n - 1
		for left < right {

			temp := value + nums[left] + nums[right]
			abs := int(math.Abs(float64(temp - target)))
			if abs < minAbs {
				minAbs = abs
				minSum = temp
			}
			switch {
			case temp > target:
				right--
			case temp < target:
				left++
			default:
				return target
			}
		}
	}
	return minSum
}
