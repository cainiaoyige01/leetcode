package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/23 9:31
 * @Description:
 */
type Foo struct {
	bar string
}

func main() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2]) //关键点是value每次都会被重用而不会开辟一块新的空间
}
