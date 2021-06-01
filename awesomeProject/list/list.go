package main

import ("container/list"
	"fmt"
)
func main(){
	link := list.New()
	// 循环插入到头部
	for i := 0; i <= 10; i++ {
		link.PushBack(i)
	}
	 n :=link.Back()//链尾
	fmt.Println("最后的值",n.Value)
	// 遍历链表 // link.Front() 链头
	for p := link.Front(); p !=  nil; p = p.Next() {
		fmt.Println("Number", p.Value)
	}

}
