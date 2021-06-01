package main

import (
	"fmt"
)
var global *int
func f() {
	var x int
	x =1
	global =&x
}
func g() {
	y := new(int)
	*y =1
}
//变量的生命周期 在函数被调用的时候，创建，函数结束之后变量销毁
func main(){
 	fmt.Println("hello ,world")
 }
//变量的作用域 全局变量（在程序结束的时候销毁） 局部变量（在函数调用结束被销毁）
// 堆和栈
//堆 分配进程内存片段 栈(堆栈) 在程序暂时创建时的局部变量