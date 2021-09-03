package main

import (
	"fmt"
)

func main() {
	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a)

	b := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(b); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, b[i])
	}
	for k,v:=range b{
		fmt.Printf("k:%d,v:%.2f\n",k,v)
	}

	//range 数组遍历
	c := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range c {//range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of c",sum)

	//多维数组
	 s := [3][4]int{
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11},   /*  第三行索引为 2 */
	}
	fmt.Println(s[0][2])
}
