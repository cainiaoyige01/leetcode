package main

/**
	给你 n 个非负整数 a1，a2，…，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
	说明：你不能倾斜容器。
	示例 1：
	输入：[1,8,6,2,5,4,8,3,7]
	输出：49
	解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
*/
import (
	"fmt"
	"math"
)

func main() {
	sum := containerWater2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(sum)
	sum1 := containerWater2([]int{4, 3, 2, 1, 4})
	fmt.Println(sum1)
	sum2 := containerWater2([]int{1, 2, 1})
	fmt.Println(sum2)

}

// containerWater 暴力解决
func containerWater(nums []int) (sum int) {
	//定义一个存水的切片
	nums2 := make([]int, len(nums)^2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := (j - i) * nums[j]
				nums2 = append(nums2, temp)
			} else {
				temp := (j - i) * nums[i]
				nums2 = append(nums2, temp)
			}
		}
	}
	sum = 0
	for _, value := range nums2 {
		if sum < value {
			sum = value
		}
	}
	return
}

// containerWater2 使用双指针来解决
func containerWater2(nums []int) (sum int) {
	//max := make([]int, len(nums)^2)
	//无需使用一个切片来进行存储 再来遍历 这样会浪费空间以及时间  可以定义一个临时变量来存储
	sum = 0
	right := len(nums) - 1
	left := 0
	for left < right {
		min := math.Min(float64(nums[left]), float64(nums[right]))
		width := right - left
		temp := int(min) * width
		if nums[left] > nums[right] {
			right--
		} else {
			left++
		}
		if temp > sum {
			sum = temp
		}

	}

	return sum
}
