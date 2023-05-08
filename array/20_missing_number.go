package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/22 15:53
 * @Description:丢失的数字
 */
func main() {
	fmt.Println(missingNumber1([]int{3, 0, 1}))
	fmt.Println(missingNumber1([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber1([]int{0}))
}

// missingNumber 先使用简单粗暴的方式来做
func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if i == num {
			continue
		} else {
		}
		return i
	}
	return len(nums)
}

// missingNumber1 利用体积差来做  还有一种方法利用哈希做
func missingNumber1(nums []int) int {
	n := len(nums)
	sum := (n * (n + 1)) / 2
	temp := 0
	for i := 0; i < n; i++ {
		temp += nums[i]
	}
	return sum - temp
}
