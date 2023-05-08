package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/20 17:44
 * @Description: 寻找峰值
 */
func main() {
	fmt.Println(findPeakElement([]int{3, 4, 3, 2, 1}))       // 1
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})) // 5 / 1
}

func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			continue
		} else {
			return i
		}
	}
	return n - 1
}
