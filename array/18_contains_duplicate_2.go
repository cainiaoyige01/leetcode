package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/21 22:57
 * @Description:存在重复元素2
 */
func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1)) // true
}

// containsNearbyDuplicate 利用哈希map来做
func containsNearbyDuplicate(nums []int, target int) bool {
	n := len(nums)
	map1 := make(map[int]int, n)
	for i, num := range nums {
		if index, ok := map1[num]; ok {
			if i-index <= target {
				return true
			}
		}
		map1[num] = i
	}
	return false
}
