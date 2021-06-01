//1.在接口中嵌入接口:
//2.在接口中嵌入结构体:
// 接口不能内嵌非接口类型，比如结构体 Interface 不能嵌入非interface的类型。
//3.在结构体中内嵌接口:
//实现接口的方法就实现了接口
//4.在结构体中嵌入结构体:
//在结构体嵌入结构体很好理解，但不能嵌入自身值类型，可以嵌入自身的指针类型即递归嵌套。
//在初始化时，内嵌结构体也进行赋值；外层结构自动获得内嵌结构体所有定义的字段和实现的方法。
package main

import (
"fmt"
)

type Writer interface {
	Write()
}

type Author struct {
	name string
	Writer
}

// 定义新结构体，重点是实现接口方法Write()
type Other struct {
	i int
}

func (a Author) Write() {
	fmt.Println(a.name, "  Write.")
}

// 新结构体Other实现接口方法Write()，也就可以初始化时赋值给Writer 接口
func (o Other) Write() {
	fmt.Println(" Other Write.")
}

func main() {

	//  方法一：Other{99}作为Writer 接口赋值
	Ao := Author{"Other", Other{99}}
	Ao.Write()

	// 方法二：简易做法，对接口使用零值，可以完成初始化
	Au := Author{name: "Hawking"}
	Au.Write()
}