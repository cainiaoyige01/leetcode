package main

import (
	"fmt"
	"sort"
)

/*
*

	给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

	示例 1：
	输入：nums = [-1,0,1,2,-1,-4]
	输出：[[-1,-1,2],[-1,0,1]]
*/
func main() {

	fmt.Println(threeSum1([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum1([]int{-1, -1, 0, 1, 1, 1, 1}))
	fmt.Println(threeSum1([]int{}))
}

// threeSum 使用双指针来解决  最大的问题是去重
func threeSum(nums []int) [][]int {
	//首先对数组进行排序
	sort.Ints(nums)
	var res [][]int
	left := 0
	right := len(nums) - 1
	for left+1 < right {
		//if nums[left] == nums[left+1] {
		//	left++
		//	continue
		//}
		//if nums[right] == nums[right-1] {
		//	right--
		//	continue
		//}
		for temp := left + 1; temp < right; temp++ {
			if nums[temp]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[left], nums[temp], nums[right]})
			}
		}
		if nums[left]+nums[left+1]+nums[right] > 0 {
			right--
		} else {
			left++
		}
	}

	return res
}

// threeSum1 使用官方的方法
func threeSum1(nums []int) (res [][]int) {
	//首先排序 有利于去重
	sort.Ints(nums)
	n := len(nums)
	for index, num := range nums {
		//如果num是大于0 就说明往后都是大于0的
		if num > 0 {
			break
		}
		//第一层遍历向前去重
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		left := index + 1
		right := n - 1
		for left < right {
			sum := num + nums[left] + nums[right]
			//在这里分为三种情况1.sum>0 2.sum<0 sum=0
			switch {
			case sum > 0:
				right--
			case sum < 0:
				left++
			default:
				res = append(res, []int{num, nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
				left++
			}
		}
	}
	return res
}
