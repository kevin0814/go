package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u User) Notify(){
	fmt.Printf("%v : %v \n",u.Name, u.Email )
}

func main(){
	u1 :=User{"golang","1035@qq.com"}
	u1.Notify()

	u2 :=User{"test","1023@qq.com"}
	u3 :=&u2
	u3.Notify()
}
