package main

import (
	"fmt"
)
func main(){
	//切片
	//var slice1 []type = make([]type, len)//slice1 := make([]type, len)
	a := [5]int{76, 77, 78, 79, 80} //指针 起始位置 和终止位置
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b) //[77, 78, 79]
	//修改切片
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]//[90, 82, 100]
	fmt.Println("array before",darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after",darr)

	//当多个片共享相同的底层数组时，每个元素所做的更改将在数组中反映出来。
	numa := [3]int{78, 79 ,80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1",numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	//切片的长度是切片中元素的数量。切片的容量是从创建切片的索引开始的底层数组中元素的数量。
	var number = make([]int,3,5) //int 指切片的类型  3指 长度 5.指最长可以达到多少
  // make 创建切片 var slice1 []type = make([]type, len)  slice1 := make([]type, len)
	printSlice(number)

	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	//append() 和 copy() 函数
	// append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
	// copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
	var numberss []int
	printSlice(numberss)

	/* 允许追加空切片 */
	numberss = append(numberss, 0)
	printSlice(numberss)

	/* 向切片添加一个元素 */
	numberss = append(numberss, 1)
	printSlice(numberss)

	/* 同时添加多个元素 */
	numberss = append(numberss, 2,3,4)
	printSlice(numberss)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers2 := make([]int, len(numberss), (cap(numberss))*2)

	/* 拷贝 numberss 的内容到 numbers2 */
	copy(numbers2,numberss)
	printSlice(numbers2)
}
func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

