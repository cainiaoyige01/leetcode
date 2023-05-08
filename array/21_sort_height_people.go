package main

import "sort"

/**
 * @Author: Niuzai
 * @Date: 2023/4/26 0:10
 * @Description:按身高排序
 */
func main() {

}

// sortPeople 使用哈希map来排序
func sortPeople(names []string, heights []int) []string {
	map1 := make(map[int]string, len(heights))
	for i := 0; i < len(heights); i++ {
		map1[heights[i]] = names[i]
	}
	sort.Ints(heights)
	j := 0
	for i := len(heights) - 1; i >= 0; i-- {
		name := map1[heights[i]]
		names[j] = name
		j++
	}
	return names
}
