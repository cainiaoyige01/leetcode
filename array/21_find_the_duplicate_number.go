package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/22 16:27
 * @Description:寻找重复数 (不适用额外的空间)
 */
func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2})) // 2
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2})) // 3
}

func findDuplicate(nums []int) int {
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] < 0 {
			return num
		}
		nums[num-1] = -nums[num-1]
	}
	return -1
}
