package main

import (
	"fmt"
)

type A struct {
	Books int
}

type B interface {
	f()
}

func (b A) f() {
	fmt.Println("A.f() ", b.Books)
}

type I int

func (i I) f() {
	fmt.Println("I.f() ", i)
}

func main() {
	var a A = A{ Books:9} //给结构体初始化并赋值 给a 变量
	a.f()

	var b B = A{Books: 99} // 接口类型可接受结构体A的值，因为结构体A实现了接口
	b.f()

	var i I = 199 // I是int类型引申出来的新类型
	i.f()

	var b2 B = I(299) // 接口类型可接受新类型I的值，因为新类型I实现了接口
	b2.f()
}