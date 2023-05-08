package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/21 14:50
 * @Description:旋转数组
 */
func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7}
	rotateArray3(ints, 3)
	fmt.Println(ints)
}

// rotateArray 每一次都把最后一位数字提上来 算是暴力解决了！
func rotateArray(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	k = k % n
	for k > 0 {
		temp := nums[n-1]
		//从后面回来可以减少一些不必要中间变量的出现
		for i := n - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = temp
		k--
	}
	return nums
}

// 这个方法太垃圾了
func rotateArray1(nums []int, k int) {
	n := len(nums)
	k = k % n
	if n == 0 {
		return
	}
	nums1 := make([]int, 0)
	nums1 = append(nums1, nums[n-k:]...)
	nums1 = append(nums1, nums[0:n-k]...)
	fmt.Println(nums1)
	//for i := 0; i < n; i++ {
	//	nums[i] = nums1[i]
	//}
	copy(nums, nums1)
}

// rotateArray 可以使用数组反转
func rotateArray2(nums []int, k int) {
	Reverse(nums)
	Reverse(nums[:k])
	Reverse(nums[k:])
}
func Reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

// rotateArray3 使用环形替换
func rotateArray3(nums []int, k int) {
	n := len(nums)
	//获取最小公约数 确定外层遍历的次数
	gcd := Gcd(k, n)
	for i := 0; i < gcd; i++ {
		index, temp := i, nums[i]
		start := index
		for true {
			cru := (k + index) % n
			nums[cru], temp = temp, nums[cru]
			index = cru
			if start == index {
				break
			}
		}
	}
}

// Gcd 找出最小公约数 这是关键点
func Gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b

}
