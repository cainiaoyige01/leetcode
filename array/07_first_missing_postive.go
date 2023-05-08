package main

import (
	"fmt"
	"sort"
)

/*
*
  - @Author: Niuzai
  - @Date: 2023/4/11 10:10
  - @Description:缺失的第一个整数

给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
进阶：你可以实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案吗？

示例 1：
输入：nums = [1,2,0]
输出：3
*/
func main() {
	fmt.Println(firstMissingPositive3([]int{3, 4, -1, 1})) // 2
	fmt.Println(firstMissingPositive3([]int{1, 2, 3, 4}))  // 5
	fmt.Println(firstMissingPositive3([]int{1, 2, 0}))     // 3
}

// firstMissingPositive3 很妙的一种思想 基于哈希表的演变
func firstMissingPositive3(nums []int) int {
	n := len(nums)
	//第一步先把数组的元素负的变为正
	for i, value := range nums {
		if value <= 0 {
			nums[i] = n + 1
		}
	}
	fmt.Println(nums)
	//然后找数字对应的长度位置变为负的
	//for _, num := range nums {
	//	//在这里还要做一个判断 如果前面的数导致了后面的值变为负 要先转为正来算
	//	num = abs(num)
	//	if num <= n {
	//		//在这里也要做判断 防止多样的数 把负值变为正了 必须确保为正转为负
	//		nums[num-1] = -abs(nums[num-1])
	//	}
	//}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n && num != 0 {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(nums int) int {
	if nums < 0 {
		return -nums
	}
	return nums
}

// firstMissingPositive 暴力解决
func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	n := len(nums) - 1
	flag := true
	temp := -1
	for i := 0; i <= n; i++ {
		if nums[i] <= 0 {
			continue
		} else if nums[i] == 1 {
			flag = false
			temp = i
			break
		} else {
			return 1
		}
	}
	if flag {
		return 1
	}
	if temp != -1 {
		for i := temp + 1; i <= n; i++ {
			//在这里还要考虑元素重复的可能行
			if nums[i-1]+1 == nums[i] || nums[i-1] == nums[i] {
				continue
			} else {
				return nums[i-1] + 1
			}

		}
	}
	return nums[n] + 1
}

// firstMissingPositive1 使用哈希表完成的  这个空间消耗大一点
func firstMissingPositive1(nums []int) int {
	n := len(nums)
	m := make(map[int]int, n)
	for i := range nums {
		m[nums[i]] = i
	}
	for i := 1; i <= n+1; i++ {

		if _, ok := m[i]; !ok {
			return i
		} else {
			continue
		}
	}
	return 1

}
