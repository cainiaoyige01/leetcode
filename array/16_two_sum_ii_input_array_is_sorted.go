package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/20 17:50
 * @Description: 两数之和||
 */
func main() {
	fmt.Println(twoSum3([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum3([]int{2, 3, 4}, 6))

}

// 使用map方法
func twoSum(nums []int, target int) []int {
	map1 := make(map[int]int, len(nums))
	sum := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		index, ok := map1[target-nums[i]]
		if ok {
			sum[0] = index
			sum[1] = i + 1
		} else {
			map1[nums[i]] = i + 1
		}
	}
	return sum
}

// twoSum3 使用左右指针
func twoSum3(nums []int, target int) []int {
	n := len(nums)
	left := 0
	right := n - 1
	sum := make([]int, 2)
	for left < right {
		temp := nums[left] + nums[right]
		switch {
		case temp > target:
			right--
		case temp < target:
			left++
		default:
			sum[0] = left + 1
			sum[1] = right + 1
			return sum
		}
	}
	return sum

}
