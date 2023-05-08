package main

import (
	"fmt"

	"github.com/valyala/bytebufferpool"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/13 16:58
 * @Description: 对象池
 */
func main() {
	object := bytebufferpool.Get()
	object.WriteString("hello world")
	fmt.Println(object)
	bytebufferpool.Put(object)
}
