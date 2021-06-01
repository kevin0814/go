package main

import "fmt"
func main(){
	sum :=0
	for i :=0; i<10;i++{
		sum +=i
	}
	fmt.Println(sum)

	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}

	//var j int
	//for ; ; j++ {
	//	if j > 10 {
	//		break
	//	}
	//	fmt.Println(j)
	//}
	//break、goto、return、panic 等语句强制退出，结束语句不会被执行。
	//var ii int
	//for ii <= 10 {
	//	ii++
	//}
	//for range 可以遍历数组、切片、字符串、map 及通道（channel）
	//for key, val := range coll {
	//	...
	//}
	//通过 for range 遍历的返回值有一定的规律：
	//数组、切片、字符串返回索引和值。
	//map 返回键和值。
	//通道（channel）只返回通道内的值。
	//遍历数组。切片
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}
	//遍历字段从 -获得字符
	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}
	//遍历 map——获得 map 的键和值
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
		}
		//对 map 遍历时，遍历输出的键值是无序的，如果需要有序的键值对输出，需要对结果进行排序。
	    //遍历通道（channel）——接收通道数据
	c := make(chan int)
	go func() {
		c <- 1
		c <- 4
		c <- 5
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	//在遍历中选择希望获得的变量
	ms := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for _, value := range ms {
		fmt.Println(value)
	}
	for key, _ := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d \n", key)
	}

	//Go语言的 for 包含初始化语句、条件表达式、结束语句，这 3 个部分均可缺省。
	//for range 支持对数组、切片、字符串、map、通道进行遍历操作。
	//在需要时，可以使用匿名变量对 for range 的变量进行选取。

	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
}

