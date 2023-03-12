package main

import (
	"fmt"
)

/**
需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

    1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
    2. 切片的长度可以改变，因此，切片是一个可变的数组。
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
    6. 如果 slice == nil，那么 len、cap 结果都等于 0。
*/

//func main() {
//	//1.声明切片
//	var s1 []int
//	if s1 == nil {
//		fmt.Println("是空")
//	} else {
//		fmt.Println("不是空")
//	}
//	// 2.:=
//	s2 := []int{}
//	// 3.make()
//	var s3 []int = make([]int, 0)
//	fmt.Println(s1, s2, s3)
//	// 4.初始化赋值
//	var s4 []int = make([]int, 0, 0)
//	fmt.Println(s4)
//	s5 := []int{1, 2, 3}
//	fmt.Println(s5)
//	// 5.从数组切片
//	arr := [5]int{1, 2, 3, 4, 5}
//	var s6 []int
//	// 前包后不包
//	s6 = arr[1:4]
//	fmt.Println(s6)
//}

// 切片初始化
// 全局：
//var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//var slice0 []int = arr[start:end]
//var slice1 []int = arr[:end]
//var slice2 []int = arr[start:]
//var slice3 []int = arr[:]
//var slice4 = arr[:len(arr)-1] //去掉切片的最后一个元素
//
////局部：
//arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
//slice5 := arr[start:end]
//slice6 := arr[:end]
//slice7 := arr[start:]
//slice8 := arr[:]
//slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素

// 代码

//var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//var slice0 []int = arr[2:8]
//var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
//var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
//var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
//var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
//func main() {
//	fmt.Printf("全局变量：arr %v\n", arr)
//	fmt.Printf("全局变量：slice0 %v\n", slice0)
//	fmt.Printf("全局变量：slice1 %v\n", slice1)
//	fmt.Printf("全局变量：slice2 %v\n", slice2)
//	fmt.Printf("全局变量：slice3 %v\n", slice3)
//	fmt.Printf("全局变量：slice4 %v\n", slice4)
//	fmt.Printf("-----------------------------------\n")
//	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
//	slice5 := arr[2:8]
//	slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
//	slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
//	slice8 := arr[0:len(arr)]  //slice := arr[:]
//	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
//	fmt.Printf("局部变量： arr2 %v\n", arr2)
//	fmt.Printf("局部变量： slice5 %v\n", slice5)
//	fmt.Printf("局部变量： slice6 %v\n", slice6)
//	fmt.Printf("局部变量： slice7 %v\n", slice7)
//	fmt.Printf("局部变量： slice8 %v\n", slice8)
//	fmt.Printf("局部变量： slice9 %v\n", slice9)
//}

// 通过make来创建切片
// var slice []type = make([]type, len)
// slice  := make([]type, len)
// slice  := make([]type, len, cap)
//var slice0 []int = make([]int, 10)
//var slice1 = make([]int, 10)
//var slice2 = make([]int, 10, 10)
//
//func main() {
//	fmt.Printf("make全局slice0 ：%v\n", slice0)
//	fmt.Printf("make全局slice1 ：%v\n", slice1)
//	fmt.Printf("make全局slice2 ：%v\n", slice2)
//	fmt.Println("--------------------------------------")
//	slice3 := make([]int, 10)
//	slice4 := make([]int, 10)
//	slice5 := make([]int, 10, 10)
//	fmt.Printf("make局部slice3 ：%v\n", slice3)
//	fmt.Printf("make局部slice4 ：%v\n", slice4)
//	fmt.Printf("make局部slice5 ：%v\n", slice5)
//	fmt.Println("--------------------------------------")
//	//	读写操作实际目标是底层数组，只需注意索引号的差别。
//	data := [...]int{0, 1, 2, 3, 4, 5}
//
//	s := data[2:4]
//	s[0] += 100
//	s[1] += 200
//
//	fmt.Println(s)
//	fmt.Println(data)
//	fmt.Println("--------------------------------------")
//	// 可直接创建 slice 对象，自动分配底层数组。
//	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
//	fmt.Println(s1, len(s1), cap(s1))
//
//	s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
//	fmt.Println(s2, len(s2), cap(s2))
//
//	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
//	fmt.Println(s3, len(s3), cap(s3))
//	fmt.Println("--------------------------------------")
//
//	//使用 make 动态创建slice，避免了数组必须用常量做长度的麻烦。还可用指针直接访问底层数组，退化成普通数组操作。
//	ss := []int{0, 1, 2, 3}
//	p := &ss[2] // *int, 获取底层数组元素指针。
//	*p += 100
//
//	fmt.Println(ss)
//	fmt.Println("--------------------------------------")
//
//	// 至于 [][]T，是指元素类型为 []T 。
//	data1 := [][]int{
//		[]int{1, 2, 3},
//		[]int{100, 200},
//		[]int{11, 22, 33, 44},
//	}
//	fmt.Println(data1)
//	fmt.Println("--------------------------------------")
//
//	// 可直接修改 struct array/slice 成员。
//	d := [5]struct {
//		x int
//	}{}
//
//	sss := d[:]
//
//	d[1].x = 10
//	sss[2].x = 20
//
//	fmt.Println(d)
//	fmt.Printf("%p, %p\n", &d, &d[0])
//	fmt.Println("--------------------------------------")
//	// 用append内置函数操作切片（切片追加）
//	var a = []int{1, 2, 3}
//	fmt.Printf("slice a : %v\n", a)
//	var b = []int{4, 5, 6}
//	fmt.Printf("slice b : %v\n", b)
//	c := append(a, b...)
//	fmt.Printf("slice c : %v\n", c)
//	d2 := append(c, 7)
//	fmt.Printf("slice d : %v\n", d2)
//	e := append(d2, 8, 9, 10)
//	fmt.Printf("slice e : %v\n", e)
//	fmt.Println("--------------------------------------")
//	//append ：向 slice 尾部添加数据，返回新的 slice 对象。
//	s11 := make([]int, 0, 5)
//	fmt.Printf("%p\n", &s11)
//
//	s22 := append(s11, 1)
//	fmt.Printf("%p\n", &s22)
//
//	fmt.Println(s11, s22)
//
//	fmt.Println("--------------------------------------")
//
//	// 超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
//	data22 := [...]int{0, 1, 2, 3, 4, 10: 0}
//	s222 := data22[:2:3]
//
//	s222 = append(s222, 100, 200) // 一次 append 两个值，超出 s.cap 限制。
//
//	fmt.Println(s222, data22)         // 重新分配底层数组，与原数组无关。
//	fmt.Println(&s222[0], &data22[0]) // 比对底层数组起始指针。
//	fmt.Println("--------------------------------------")
//
//	//[0 1 100 200] [0 1 2 3 4 0 0 0 0 0 0]
//	//0xc4200160f0 0xc420070060
//	/**
//	从输出结果可以看出，append 后的 s 重新分配了底层数组，并复制数据。如果只追加一个值，则不会超过 s.cap 限制，也就不会重新分配。
//	通常以 2 倍容量重新分配底层数组。在大批量添加数据时，建议一次性分配足够大的空间，以减少内存分配和数据复制开销。
//	或初始化足够长的 len 属性，
//	改用索引号进行操作。及时释放不再使用的 slice 对象，避免持有过期数组，造成 GC 无法回收。
//	*/
//
//	// slice中cap重新分配规律
//	s44 := make([]int, 0, 1)
//	c22 := cap(s44)
//
//	for i := 0; i < 50; i++ {
//		s44 = append(s44, i)
//		if n := cap(s44); n > c22 {
//			fmt.Printf("cap: %d -> %d\n", c22, n)
//			c22 = n
//		}
//	}
//	fmt.Println("--------------------------------------")
//}

// 切片拷贝
//func main() {
//
//	s1 := []int{1, 2, 3, 4, 5}
//	fmt.Printf("slice s1 : %v\n", s1)
//	s2 := make([]int, 10)
//	fmt.Printf("slice s2 : %v\n", s2)
//	copy(s2, s1)
//	fmt.Printf("copied slice s1 : %v\n", s1)
//	fmt.Printf("copied slice s2 : %v\n", s2)
//	s3 := []int{1, 2, 3}
//	fmt.Printf("slice s3 : %v\n", s3)
//	s3 = append(s3, s2...)
//	fmt.Printf("appended slice s3 : %v\n", s3)
//	s3 = append(s3, 4, 5, 6)
//	fmt.Printf("last slice s3 : %v\n", s3)
//	fmt.Println("--------------------------------------")
//	// copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。
//	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	fmt.Println("array data : ", data)
//	s11 := data[8:]
//	s22 := data[:5]
//	fmt.Printf("slice s1 : %v\n", s11)
//	fmt.Printf("slice s2 : %v\n", s22)
//	copy(s22, s11)
//	fmt.Printf("copied slice s1 : %v\n", s11)
//	fmt.Printf("copied slice s2 : %v\n", s22)
//	fmt.Println("last array data : ", data)
//	fmt.Println("--------------------------------------")
//	// 应及时将所需数据 copy 到较小的 slice，以便释放超大号底层数组内存。
//
//}

// slice遍历：
//func main() {
//
//	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	slice := data[:]
//	for index, value := range slice {
//		fmt.Printf("inde : %v , value : %v\n", index, value)
//	}
//	fmt.Println("--------------------------------------")
//	// 切片resize（调整大小）
//	var a = []int{1, 3, 4, 5}
//	fmt.Printf("slice a : %v , len(a) : %v\n", a, len(a))
//	b := a[1:2]
//	fmt.Printf("slice b : %v , len(b) : %v\n", b, len(b))
//	c := b[0:3]
//	fmt.Printf("slice c : %v , len(c) : %v\n", c, len(c))
//	fmt.Println("--------------------------------------")
//}

// 字符串和切片（string and slice）
// string底层就是一个byte的数组，因此，也可以进行切片操作。

func main() {
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)

	//string本身是不可变的，因此要改变string中字符。需要如下操作：
	//英文字符串：
	str1 := "Hello world"
	s := []byte(str1) //中文字符需要用[]rune(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str1 = string(s)
	fmt.Println(str1)

	// 含有中文字符串：
	str3 := "你好，世界！hello world！"
	s3 := []rune(str3)
	s3[3] = '够'
	s3[4] = '浪'
	s3[12] = 'g'
	s3 = s3[:14]
	str3 = string(s3)
	fmt.Println(str3)
	/**
	golang slice data[:6:8] 两个冒号的理解

	常规slice , data[6:8]，从第6位到第8位（返回6， 7），长度len为2， 最大可扩充长度cap为4（6-9）

	另一种写法： data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8
	*/

	// a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := slice[6:8]
	fmt.Println(d1, len(d1), cap(d1))
	d2 := slice[:6:8]
	fmt.Println(d2, len(d2), cap(d2))

	// 数组or切片转字符串：
	//strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)
}
