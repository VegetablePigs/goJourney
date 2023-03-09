package main

import "fmt"

func main() {
	//var cat int = 10
	//var str = "雷电将军"
	//
	//// 指针赋值
	//prt := &str
	//fmt.Printf("%p,%p \n", &cat, &str)
	//// 取值
	//fmt.Printf(*prt)

	// 指针与变量
	//var room int = 10 // room房间 里面放的 变量10
	//var ptr = &room   // 门牌号px  指针  0xc00000a0a8
	//
	//fmt.Printf("%p\n", &room) // 变量的内存地址 0xc00000a0a8
	//
	//fmt.Printf("%T, %p\n", ptr, ptr) // *int, 0xc00000a0a8
	//
	//fmt.Println("指针地址", ptr)      // 0xc00000a0a8
	//fmt.Println("指针地址代表的值", *ptr) // 10

	// 利用指针修改值
	var num = 10
	modifyFromPoint(num)
	fmt.Println("未使用指针，方法外", num)

	var num2 = 22
	newModifyFromPoint(&num2) // 传入指针
	fmt.Println("使用指针 方法外", num2)
}

func newModifyFromPoint(ptr *int) {
	// 使用指针
	*ptr = 1000 // 修改指针地址指向的值
	fmt.Println("使用指针，方法内:", *ptr)
}

func modifyFromPoint(num int) {
	// 未使用指针
	num = 10000
	fmt.Println("未使用指针，方法内:", num)
}
