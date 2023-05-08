package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/9 23:10
 * @Description:
 */
func main() {
	//注意Student是实现了People的接口了
	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	//有一个点：当接口的动态值以及动态类型都为nil的时候 接口才为nil！在这里动态值为nil 但是动态类型不为nil
	var p People = s // p = (*main.Student)(nil)
	fmt.Printf("%#v\n", p)
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}
