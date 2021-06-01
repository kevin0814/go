package main

import "fmt"
/*    1. 数组：是同一种数据类型的固定长度的序列。
    2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
    3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
    4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
    for i := 0; i < len(a); i++ {
    }
    for index, v := range a {
    }
    5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
    6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
    7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
    8.指针数组 [n]*T，数组指针 *[n]T。*/
func main(){
	//一维数组 var 数组变量名 [元素数量]Type
	var arr_1 [5]int =[5]int{1,2,4,5,7}
	fmt.Println(arr_1)
for i,v :=range arr_1 {
		fmt.Printf("%d %d \n",i,v)
	}
for _,v := range arr_1 {
	fmt.Printf("%d \n",v)
}
var r [3]int =[3]int {1,4}
 fmt.Println(r[2])
 q := [...] int {1, 2, 3}
fmt.Printf("%T \n",q)

 m :=[3]int{1,3,5}
 m =[3]int{1,2,43}
 fmt.Println(m)

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"

	var team [4] string
	team[2] ="hamm"
	/*team[5] ="soldier"
	team[7] ="mum"*/
	 for k,v :=range team{
	 	fmt.Println(k,v)
	 }
   //多维数组 命名 方式
	array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array)
	//指定数组索引 1和3的元素
	array1 := [4][2]int{2: {20, 21}, 1: {40, 41}}
    fmt.Println(array1)
	// 声明并初始化数组中指定的元素
   var array5[4][2]int = [4][2]int{1: {0: 20}, 3: {1: 41}}
   //等价于 array5 :=[4][2]int{1:{0:2},3:{3:10}}
   fmt.Println(array5)

   var arr =[5] int{1,2,3,4,5}
   modifyArr(&arr)
   fmt.Println(arr)
}

func modifyArr(a *[5] int){
	a[1]=10
}


