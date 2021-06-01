package main

import (
	"fmt"
	"math"
)
func main(){
	//函数的基本组成为：关键字 func、函数名、参数列表、返回值、函数体和返回语句
	//Go语言里面拥三种类型的函数：
	//普通的带有名字的函数
	//匿名函数或者 lambda 函数
	//方法
	//普通函数声明(定义)
	//形式参数列表描述了函数的参数名以及参数类型，这些参数作为局部变量
	//func 函数名(形式参数列表)(返回值列表){
	//	函数体
	//}
	fmt.Println(hypot(3,4)) // "5"
	//result := add(1,1)
	//	//println(result)

   //	Go语言函数变量——把函数作为值保存到变量中
   var  f func()
	f=fire
	f()
	func(data int) {
		fmt.Println("hello", data)
	}(100)

	myfunc(2, 3, 4)
	myfunc(22, 33, 44,55)
}
func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
func fire() {
	fmt.Println("fire")
}
//可变参数，没有固定参数个数
func myfunc(args ...int) { //...type   []type 切片
	for _, arg := range args {
		fmt.Println(arg)
	}
}

