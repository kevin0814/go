package main

import (
	"fmt"
)

func main (){
	var  a =[3] int{1,2,3}
	fmt.Println(a,a[2:3])  
	

	var sli_1 [] int //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v \n",len(sli_1),cap(sli_1),sli_1)
	var highRiseBuilding [30]int

	for i :=0;i<30; i++ { //赋值
		highRiseBuilding[i] =i+1
	}
	fmt.Println(highRiseBuilding)
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])

	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])
	//数组
	a =[3]int{1,3,5}
	//表示原有切片
	a1 :=[3]int{1,2,5}
	fmt.Println(a1[0:0])
	fmt.Println(a1[:])
	fmt.Println(a1)
	fmt.Println(a)

	//声明切片 int 类型
	var names [] int
	//声明切片字符串类型
	var str [] string
	//元素类型  string “”``  bool true false number  整型 int8 int16 int32 int64   浮点型 float32 float64 复数
	// 声明一个空切片
	var numListEmpty = []int{}

	fmt.Println(names,str,numListEmpty)
	fmt.Println(len(names), len(str), len(numListEmpty)) //len 获取切片大小
	// 切片判定空的结果
	fmt.Println(names == nil)
	fmt.Println(str == nil)
	fmt.Println(numListEmpty == nil)

	 //使用 make() 函数构造切片
	 //make( []Type, size, cap ) 其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题
	s := make([]int, 2) //构造 切片
	l := make([]int, 2, 10)
	fmt.Println(s, l)
	fmt.Println(len(s), len(l))

	var sli_2 [2] int      //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_2),cap(sli_2),sli_2)
	var sli_12 = [] int {} //空切片
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_12),cap(sli_12),sli_12)

	var sli_3 = [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_3),cap(sli_3),sli_3)

	sli_4 := [] int {1, 2, 3, 4, 5} //切片 片段
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_4),cap(sli_4),sli_4)

	var sli_5 [] int = make([] int, 5, 8)//构造切片
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_5),cap(sli_5),sli_5)

	sli_6 := make([] int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli_6),cap(sli_6),sli_6)

	//截取切片
	sli := [] int {1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)

	fmt.Println("sli[1] ==", sli[1])//[2]
	fmt.Println("sli[:] ==", sli[:]) //[1,2,3,4,5,6]
	fmt.Println("sli[1:] ==", sli[1:])//[2,3,4,5,6]
	fmt.Println("sli[:4] ==", sli[:4])//[1,2,3,4]

	fmt.Println("sli[0:3] ==", sli[0:3])//[1,2,3]
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[0:3]),cap(sli[0:3]),sli[0:3])

	fmt.Println("sli[0:3:4] ==", sli[0:3:4])
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli[0:3:4]),cap(sli[0:3:4]),sli[0:3:4])

	fmt.Println("---追加切片\n")
	//追加切片
	sli2 := [] int {4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli2),cap(sli2),sli2)

	sli2 = append(sli2, 7)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli2),cap(sli2),sli2)

	sli2 = append(sli2, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli2),cap(sli2),sli2)

	sli2 = append(sli2, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli2),cap(sli2),sli2)

	sli2 = append(sli2, 10)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli2),cap(sli2),sli2)
	//删除切片
	sli3 := [] int {1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli3),cap(sli3),sli3)

	//删除尾部 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli3[:len(sli3)-2]),cap(sli3[:len(sli3)-2]),sli3[:len(sli3)-2])

	//删除开头 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli3[2:]),cap(sli3[2:]),sli3[2:])

	//删除中间 2 个元素
	sli = append(sli3[:3], sli3[3+2:]...)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(sli3),cap(sli3),sli3)
}
