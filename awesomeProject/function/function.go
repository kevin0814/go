package main

import (
	"fmt"
	"time"
)

type funcType func(time.Time)     // 定义函数类型funcType

func main() {
	s := make([]int, 10, 10)
	ss :=make(map[string]int)
	var b []byte
	b = append(b, "bar111"...)
	fmt.Printf("%s \n",b)
	//new(T)内置函数在运行时为该类型的变量分配内存，并返回指向它的类型* T的值。 并对变量初始化
	fmt.Println(ss)
	fmt.Println(s)
	f := func(t time.Time) time.Time { return t } // 方式一：直接赋值给变量
	fmt.Println(f(time.Now()))

	var timer funcType = CurrentTime // 方式二：定义函数类型funcType变量timer
	timer(time.Now())

	funcType(CurrentTime)(time.Now())  // 先把CurrentTime函数转为funcType类型，然后传入参数调用

//	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
//	var s = make([]int, 6)
	var bb = make([]byte, 5)
	//n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
	//n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
	n3 := copy(bb, "Hello, World!")  // n3 == 5, b == []byte("Hello")
	fmt.Printf("%s",bb)
	fmt.Println(n3)
	// 这种处理方式在Go 中比较常见
}
//在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）等这样的引用类型都是默认使用引用传递。

func CurrentTime(start time.Time) {
	fmt.Println(start)
}