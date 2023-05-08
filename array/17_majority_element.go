package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/20 18:50
 * @Description: 多数元素
 */
func main() {
	fmt.Println(majorityElement([]int{6, 5, 5}))
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
func majorityElement(nums []int) int {
	//创建切片
	map1 := make(map[int]int, len(nums))
	for _, value := range nums {
		map1[value]++
	}
	fmt.Println(map1)
	sum := 0
	temp := 0
	for i := range map1 {
		i2 := map1[i]
		if i2 > sum {
			sum = i2
			temp = i
		}
	}
	return temp
}
