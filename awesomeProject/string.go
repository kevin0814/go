package main

import (
	"fmt"
	"math"
)

func main(){
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte UTF8编码下一个中文汉字由3~4个字节组成
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r) //
	}
	fmt.Println()

	//要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))//转换为字符串

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))

	//类型转换 T(表达式) T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}