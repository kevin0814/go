package main

import (
	"fmt"
)

func main() {
	// for 循环  没有while
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d",i)
	}

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	//循环使用 加跳出循环 break
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
	//for 遍历
	//for i,x:= range numbers{
	//	fmt.Printf("第 %d 位 x的值 =%d \n",i,x)
	//}

	//continue：跳出一次循环。continue语句用于跳过for循环的当前迭代
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	var m int = 10
//loop 跳出循环 跳过指定循环，循环不会终止
	/* 循环 */
LOOP: for m < 20 {
	if m == 15 {
		/* 跳过迭代 */
		m = m + 1
		goto LOOP
	}
	fmt.Printf("a的值为 : %d\n",m)
	m++
}
}
