package main

import (
"fmt"
)

type A struct {
	Face int
}
type Aa A // 自定义新类型Aa，没有基础类型A的方法

func (a A) f() {
	fmt.Println("hi ", a.Face)
}
func (a Aa) f() {
	fmt.Println("hi ", a.Face)
}

func main() {
	var s A = A{ Face: 9 } //A 指类型
	s.f()

	var sa Aa = Aa{ Face: 9 }
	sa.f()
	ss :=func(x, y int) int { return x + y } (3, 4) //可以直接对匿名函数进行调用，注意匿名函数的最后面加上了括号并填入了参数值，如果没有参数，也需要加上空括号，代表直接调用
	fmt.Println(ss)

	func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
		fmt.Println(sum)
	}()

}
