package main

import "fmt"

// demo 值交换
func main() {
	// 第一种
	//var (
	//	a = 100
	//	b = 200
	//	c int
	//)
	//
	//c = b
	//b = a
	//a = c
	//
	//fmt.Printf("a=%d,b=%d", a, b)

	// 第二种
	//a := 100
	//b := 200
	//
	//a = a ^ b
	//b = b ^ a
	//a = a ^ b
	//fmt.Printf("a=%d,b=%d", a, b)

	// 第三种
	a := 100
	b := 200
	b, a = a, b
	fmt.Printf("a=%d,b=%d", a, b)

}
