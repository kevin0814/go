package main

import (
	"fmt"
	"time"
	)

const pi  =3.14159
const name string ="zhangsan"
const  (
	e= 2.12312
	pis =3.132323
)
const (
	a = 1
	b
	c = 3
	d
)
//const IPv4Len =4
//
//func parseIpv4(s string) IP {
//	var p [IPv4Len]byte
//	return p
//}
func main(){
	fmt.Println("here")
    fmt.Println(name)
     ll :=name + " world"
   fmt.Println(ll)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	fmt.Println(a, b, c, d)
}
