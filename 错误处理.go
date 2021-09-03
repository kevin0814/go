package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./test1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//step3：关闭文件
	defer f.Close()

	//step2：读取数据
	bs := make([]byte,10,20)
	n,err :=f.Read(bs)
	fmt.Println(err) //<nil>
	fmt.Println(n) //4
	fmt.Println(string(bs)) //[97 98 99 100]

	//ff, err :=os.Stat("./test1.txt")
	//fmt.Printf("%T\n",ff)
	//fmt.Println(ff.Mode())
	//fmt.Println(ff.Name())
	//fmt.Println(ff.Size())
	//fmt.Println(n)

	//根据f进行文件的读或写
	fmt.Println(f.Name(), "opened successfully")
}
