package main

import ("container/list"
	"fmt"
)

func main(){
	var vaLis list.List
	//双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。
	//l := list.New()
	vaLis.PushBack("fist")
	vaLis.PushFront(67)
	//
	//l := list.New()
	//// 尾部添加
	//l.PushBack("canon")
	//// 头部添加
	//l.PushFront(67)
	//// 尾部添加后保存元素句柄
	//element := l.PushBack("fist")
	//// 在fist之后添加high
	//l.InsertAfter("high", element)
	//// 在fist之前添加noon
	//l.InsertBefore("noon", element)
	//// 使用
	//l.Remove(element)

	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	l.PushBack("keccc")
	// 头部添加
	l.PushFront(67)
	for i :=l.Front();i!=nil;i= i.Next(){
		fmt.Println(i.Value,i.Prev())
	}
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	ll := list.New()
	// 尾部添加
	ll.PushBack("canon")
	// 头部添加
	ll.PushFront(67)
	// 尾部添加后保存元素句柄
	element := ll.PushBack("fist")
	// 在fist之后添加high
	ll.InsertAfter("high", element)
	// 在fist之前添加noon
	ll.InsertBefore("noon", element)
	// 使用
	ll.Remove(element)
}