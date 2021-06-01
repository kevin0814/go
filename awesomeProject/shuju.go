package main

import (
	"fmt"
)
func main(){
	//变量声明规则  var 变量名 变量类型
	//名字 := 表达式(简短格式)
	//定义变量，同时显式初始化。
	//不能提供数据类型。
	//只能用在函数内部。
	var name int
	fmt.Println(name)
	//字符串
	str :="hello"
	fmt.Println(str)
	stre :="Beginning of the string" +
		"second"
	stre +="world"
	fmt.Println(stre)
	sre :="hello"
	fmt.Println(sre[1])
	fmt.Println(len(str))
	a := 5.0
	b := int(a)
	fmt.Println(b)
    //	new(类型) 创建指针的另一种方法
    //new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。
  //常量 const name [type] = value
	//显式类型定义： const b string = "abc"
	//隐式类型定义： const b = "abc"
	//分别是标识符（identifier）、关键字（keyword）、操作符（operator）、分隔符（delimiter）、字面量（literal）
   //var 数组变量名 [元素数量]Type 数组声明
   //多维数组 var array_name [size1][size2]...[sizen] array_type
   //array = [4][2]int{1: {0: 20}, 3: {1: 41}}
   //切片  slice[开始位置:结束位置]
   //声明切片 var name []Type
	//使用 make() 函数构造切片 make( []Type, size, cap )
	//复制切片 copy( destSlice, srcSlice []T) int
	//var sliceName [][]...[]sliceType 多维切片
	 //map 映射 键值对 var mapname map[keytype]valuetype
	 //初始化列表 container/list
	//变量名 := list.New()
	// var 变量名 list.List

}