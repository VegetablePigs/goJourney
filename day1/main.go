package main

import (
	"fmt"
	"net"
	"strconv"
)

// 声明变量
var age int16 = 9
var name = "张三"
var level float32 = 1.0

// 批量声明变量
var (
	a int
	b string
	c float32
)

// 简短格式 可以省略 var
//i := 1

func main() {
	//fmt.Println(name)
	//fmt.Printf("%T", level) // 推断类型
	//%d整数占位符，%s 字符串占位符，%f 浮点数占位符(默认精度为6)
	//fmt.Printf("a=%d , b=%s , c=%f", a, b, c)

	// 简短格式 可以省略 var 只能定义在方法内部
	/**
	简短格式限制：
	1.只能定义在方法内部
	2.不能提供数据类型
	3.创建时要显示的初始化
	*/
	//i := 1 // 方法内部定义的变量必须要使用
	//fmt.Print(i)

	// 变量初始化不能重复定义
	//var lev = 1
	////lev := 2  报错
	//fmt.Print(lev)

	// 多重赋值 只要左边有一个值重新定义了（conn1）就可以重复赋值
	// 在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，即便其他变量名可能是重复声明的，编译器也不会报错
	//var conn net.Conn
	//var err error
	//conn, err := net.Dial("tcp", "127.0.0.1:8088")
	//conn1, err := net.Dial("tcp", "127.0.0.1:8088")
	//fmt.Println(conn)
	//fmt.Println(err)
	//fmt.Println(conn1)

	// 匿名变量
	//如果不想接收err的值，那么可以使用_表示，这就是匿名变量
	/**
	匿名变量以“_”下划线表示
	匿名变量不占用命名空间，也不会分配内存。匿名变量可以重复声明使用

	”_”木身就是一个特殊的标识符，被称为空白标识符。它可以像其他标识符那样用于变量的声明或赋值(任何类型都可以赋值给它），
	但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。
	*/
	//conn, err := net.Dial("tcp", "127.0.0.1:8088")
	//匿名变量不可以直接开头
	// - := 1
	conn, _ := net.Dial("tcp", "127.0.0.1:8088")
	fmt.Println(conn)
	//fmt.Println(err)

	// 字符串与其他类型的转换
	// str 转 int
	newStr1 := "1"
	intValue, _ := strconv.Atoi(newStr1)
	fmt.Printf("%T,%d\n", intValue, intValue) // int,1

	// int 转 str
	intValue2 := 1
	strValue := strconv.Itoa(intValue2)
	fmt.Printf("%T, %s\n", strValue, strValue)

	// str 转  float
	string3 := "3.1415926"
	f, _ := strconv.ParseFloat(string3, 32)
	fmt.Printf("%T, %f\n", f, f) // float64, 3.141593
	//float 转 string
	floatValue := 3.1415926
	//4个参数，1：要转换的浮点数 2. 格式标记（b、e、E、f、g、G）
	//3. 精度 4. 指定浮点类型（32:float32、64:float64）
	// 格式标记：
	// ‘b’ (-ddddp±ddd，二进制指数)
	// ‘e’ (-d.dddde±dd，十进制指数)
	// ‘E’ (-d.ddddE±dd，十进制指数)
	// ‘f’ (-ddd.dddd，没有指数)
	// ‘g’ (‘e’:大指数，‘f’:其它情况)
	// ‘G’ (‘E’:大指数，‘f’:其它情况)
	//
	// 如果格式标记为 ‘e’，‘E’和’f’，则 prec 表示小数点后的数字位数
	// 如果格式标记为 ‘g’，‘G’，则 prec 表示总的数字位数（整数部分+小数部分）
	formatFloat := strconv.FormatFloat(floatValue, 'f', 2, 64)
	fmt.Printf("%T,%s", formatFloat, formatFloat)
}
