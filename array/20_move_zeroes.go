package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/22 16:15
 * @Description:移动零
 */
func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0})
}

// moveZeroes 使用快慢指针来做  还有一种方法不覆盖 在后面补零
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	n := len(nums)
	for fast < n {
		if nums[fast] == 0 {
			fast++
		} else {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
			fast++
		}
	}
	fmt.Println(nums)
}
