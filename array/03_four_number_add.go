package main

import (
	"fmt"
	"sort"
)

/*
*给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：答案中不可以包含重复的四元组。
示例 1：
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
*/
//做法应该是双层嵌套循环 利用指针
func main() {
	fmt.Println(fourSum1([]int{1, 0, -1, 0, -2, 2}, 0)) // [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
	fmt.Println(fourSum1([]int{-3, 0, 0, 1, 2}, 0))     // [[-3 0 1 2]]
	fmt.Println(fourSum1([]int{0, 0, 0, 0}, 0))         // [[0 0 0 0]]

}

// fourSum1 去重的特色 可以学一下
func fourSum1(nums []int, target int) [][]int {
	sort.Ints(nums)
	var fourNumber [][]int
	n := len(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := n - 1
			for left < right {
				temp := nums[i] + nums[j] + nums[left] + nums[right]
				switch {
				case temp > target:
					right--
				case temp < target:
					left++
				default:
					fourNumber = append(fourNumber, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right++
					}
					left++
					right--
				}
			}
		}
	}

	return fourNumber
}

// fourSum 自己利用map的特性完成去重
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var fourNumber [][]int
	n := len(nums)
	m := make(map[string][]int, n)
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			left := j + 1
			right := n - 1
			for left < right {
				temp := nums[i] + nums[j] + nums[left] + nums[right]
				switch {
				case temp > target:
					right--
				case temp < target:
					left++
				default:
					str := strNums([]int{nums[i], nums[j], nums[left], nums[right]})
					m[str] = []int{nums[i], nums[j], nums[left], nums[right]}
					left++
					right--
				}
			}
		}
	}
	for _, value := range m {
		//fmt.Println(value)
		fourNumber = append(fourNumber, value)
	}
	return fourNumber
}
func strNums(nums []int) string {
	var sprintf = ""
	for _, v := range nums {
		sprintf += fmt.Sprintf("%d", v)
	}
	return sprintf
	//str := ""
	//for _, num := range nums {
	//	str += fmt.Sprintf("%d_", num)
	//}
	//return str

}
