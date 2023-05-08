/**
 * @Author: _niuzai
 * @Date:   2023/5/9 0:24
 * @Description:init函数的初始化
 *
 */
package main

import "fmt"

func main() {
	/**
	init() 函数是 Go 程序初始化的一部分。Go 程序初始化先于 main 函数，由 runtime 初始化每个导入的包，
	初始化顺序不是按照从上到下的导入顺序，而是按照解析的依赖关系，没有依赖的包最先初始化。

	每个包首先初始化包作用域的常量和变量（常量优先于变量），然后执行包的 init() 函数。同一个包，
	甚至是同一个源文件可以有多个 init() 函数。init() 函数没有入参和返回值，不能被其他函数调用，
	同一个包内多个 init() 函数的执行顺序不作保证。

	一句话总结： import –> const –> var –> init() –> main()
	*/
	fmt.Println("main:", a)
}
func init() {
	fmt.Println("init1:", a)
}

func init() {
	fmt.Println("init2:", a)
}

var a = 10

const b = 100
