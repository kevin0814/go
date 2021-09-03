package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a chan int
	// 通道 申明 a := make(chan int)   a := make(chan int)
	if a == nil {
		fmt.Println("channel 是 nil 的, 不能使用，需要先创建通道。。")
		a = make(chan int)
		fmt.Printf("数据类型是： %T", a)
	}

	ch1 := make(chan int)
	fmt.Printf("%T,%p\n",ch1,ch1)
   if reflect.TypeOf(a)==reflect.TypeOf(ch1) {
   	fmt.Println("同一个通道")
   	fmt.Println(reflect.TypeOf(a))
   }else{
   	fmt.Println("不是同一个通道")
   }
	test1(ch1)
}

func test1(ch chan int){
	fmt.Printf("%T,%p\n",ch,ch)
}
