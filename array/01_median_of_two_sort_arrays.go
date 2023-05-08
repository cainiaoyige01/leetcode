package main

import "fmt"

/**
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
*/

func main() {
	fmt.Println(findMedianSortArrays1([]int{}, []int{}))                  // OJ 无此类异常 case
	fmt.Println(findMedianSortArrays1([]int{1, 3, 5}, []int{2, 4, 6}))    // 3.5
	fmt.Println(findMedianSortArrays1([]int{1, 3, 5, 7}, []int{2, 4, 6})) // 4
	fmt.Println(findMedianSortArrays1([]int{1, 3, 5}, []int{2, 4, 6}))    // 3.5
}

// findMedianSortArrays1
func findMedianSortArrays1(nums1 []int, nums2 []int) float64 {
	//对数组进行合并
	sortArray := merge(nums1, nums2)
	fmt.Println(sortArray)
	length := len(sortArray)
	if length == 0 {
		return -1
	}
	if length%2 == 1 {
		fmt.Printf("%.5f\n", float64(sortArray[length/2]))
		return float64(sortArray[length/2])
	} else {
		fmt.Printf("%.5f\n", float64(sortArray[length/2]+sortArray[length/2-1])/2)
		return float64(sortArray[length/2]+sortArray[length/2-1]) / 2
	}
}

// merge 这里的核心思想是利用append进行数据的拼接
func merge(nums1 []int, nums2 []int) []int {
	//获取数组的长度
	n1, n2 := len(nums1), len(nums2)
	//对切片进行初始化
	targetArray := make([]int, 0, n1+n2)
	i, j := 0, 0
	for i < n1 && j < n2 {
		//这里有三种情况 1.n1的元素比n2的大 2.n1的元素比n2的小 3.n1的元素比n2的相等
		switch {
		case nums1[i] < nums2[j]:
			targetArray = append(targetArray, nums1[i])
			i++
		case nums1[i] > nums2[j]:
			targetArray = append(targetArray, nums2[j])
			j++
		default:
			targetArray = append(targetArray, nums1[i], nums2[j])
			i++
			j++
		}
	}
	// 在进行数组剩下的元素拼接到目标数组中去
	if i < n1 {
		targetArray = append(targetArray, nums1[i:]...)
	}
	if j < n2 {
		targetArray = append(targetArray, nums2[j:]...)
	}
	return targetArray

}
