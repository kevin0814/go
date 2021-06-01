package main
import (
	"bytes"
	"fmt"
    "math"
	"strings"
)
// 将NewInt定义为int类型
type NewInt int
// 将int取一个别名叫IntAlias
type IntAlias = int

func main() {
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)
	const e = .718128 // 0.718128
	const f = 1.     // 1
	fmt.Println(e);fmt.Println(f)
	fmt.Printf("%f \n",math.Pi)
	fmt.Printf("%.2f\n",math.Pi) //获取两位小数点
    //复数 由实数+虚数 i 虚数单位  RE + IMi，
    //var name complex128 = complex(x,y)
    var x complex128 =complex(1,2)
    var y complex128 =complex(3,4)
    fmt.Println(x* y)
    var c  int =10
    if 'a'<=c && c<='z' ||
    	'A'<=c && c<='Z' ||
    	'0'<=c && c<='9'{
    	println(111)
	}
	var str ="C语音中文网\n Go语音教程"
	fmt.Println(str)
	fmt.Println(str[0],str[5-1],str[len(str)-1])
	//字符串拼接
	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”
const	strs = `第一行
第二行
第三行
\r\n
`
	fmt.Println(strs)
//字符串拼接
s1 :=fmt.Sprintf("%s:%d","127.0.0.1",8080)
fmt.Println(s1)
//字符串拼接2
s2 := strings.Join([]string{"hello ","world"},"")
fmt.Println(s2)
//字符串拼接3
s3 :=bytes.Buffer{}
s3.WriteString("hello ")
s3.WriteString("world s3")
fmt.Println(s3.String())
//拼接字符串4
s4 :="hello " + "world"
fmt.Println(s4)
//获取字符串长度
tip1 := "genji is a ninja"
fmt.Println(tip1)

tip2 :="忍者"
fmt.Println(len(tip2))

var ch int ='\u0041'
var ch2 int ='\u03B2'
var ch3 int ='\U00101234'
fmt.Printf("%d -%d -%d",ch,ch2,ch3)
fmt.Printf("%c - %c - %c",ch,ch2,ch3)
fmt.Printf("%X -%X -%X",ch,ch2,ch3)
fmt.Printf("%U - %U -%U \n",ch,ch2,ch3)

 //输出各数值范围
 //fmt.Println("int8 range:",math.MinInt8,math.MaxInt8)
 //fmt.Println("int16 range:",math.MinInt16,math.MaxInt16)
 //fmt.Println("int32 range",math.MinInt32,math.MaxInt32)
 //fmt.Println("int64 range",math.MinInt64,math.MaxInt64)
  var cat int =1
  var st string ="banana"
  fmt.Printf("%p %p",&cat,&st)
}
