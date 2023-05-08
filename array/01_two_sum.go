package main

import (
	"fmt"
)

/*
*

	两数相加
	给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
	你可以按任意顺序返回答案
	输入：nums = [2,7,11,15], target = 9
	输出：[0,1]
	解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

*
*/
func main() {
	fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum2([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum2([]int{3, 3}, 6))         // [0,1]
}

// twoSum1 第一种方法进行暴力解题
func twoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// twoSum2 第二种方式利用哈希表索引 用数组的值作为map键 下标作为map值
func twoSum2(nums []int, target int) []int {
	//要初始化map
	mapSum := make(map[int]int, len(nums))
	//遍历数组nums 进行添加元素以及检查元素
	for index, value := range nums {
		i, ok := mapSum[target-value]
		if ok && i != index {
			return []int{i, index}
		}
		//注意 这个不能放在前面  防止有相同的元素
		mapSum[value] = index
	}
	return nil
}
