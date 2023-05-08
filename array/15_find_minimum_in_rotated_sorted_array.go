package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/18 10:47
 * @Description:寻找旋转排序数组中的最小值
 */
func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}

// findMin 寻找最小值 使用二分原理!利用好升序这个条件 参数不给参考值 就自定义一个
func findMin(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return nums[0]
	}
	if nums[n-1] > nums[0] {
		return nums[0]
	}
	left := 0
	right := n - 1
	target := nums[0]
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target { //中间值比目标值要大 说明前面是有序的

			left = mid + 1
		} else if nums[mid] < target {
			right = mid
		}
	}
	return nums[right]
}
