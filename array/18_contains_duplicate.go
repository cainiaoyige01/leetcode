package main

import (
	"fmt"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/21 22:46
 * @Description:存在重复元素
 */
func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1})) // false
}

func containsDuplicate(nums []int) bool {
	n := len(nums)
	map1 := make(map[int]int, n)
	for i, num := range nums {
		if _, ok := map1[num]; ok {
			return true
		}
		map1[num] = i
	}
	return false
}
