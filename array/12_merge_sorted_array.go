package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/16 9:43
 * @Description:合并两个有序数组
 */
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge1(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

// merge1 合并数组
func merge1(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		nums1[m] = nums2[i]
		m++
	}
	sort.Ints(nums1)
}
