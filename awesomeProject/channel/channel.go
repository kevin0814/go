package main

func main(){
	var channel chan int = make(chan int)
	// Go中通道可以是发送（send）、接收（receive）、同时发送（send）和接收（receive）。
	// 定义接收的通道
	//receive_only := make (<-chan int)
	// 定义发送的通道
	//send_only := make (chan<- int)
	// 可同时发送接收
	//send_receive := make (chan int)
	// chan<- 表示数据进入通道，要把数据写进通道，对于调用者就是发送。
	// <-chan 表示数据从通道出来，对于调用者就是得到通道的数据，当然就是接收。
}
