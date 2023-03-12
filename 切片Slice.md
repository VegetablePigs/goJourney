## 切片Slice
需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

    1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
    2. 切片的长度可以改变，因此，切片是一个可变的数组。
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
    6. 如果 slice == nil，那么 len、cap 结果都等于 0。

```go
package main

import (
	"fmt"
)

/**
Golang Array和以往认知的数组有很大不同。
	1. 数组：是同一种数据类型的固定长度的序列。
    2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
    3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
    4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
    for i := 0; i < len(a); i++ {
    }
    for index, v := range a {
    }
    5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
    6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
    7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
    8.指针数组 [n]*T，数组指针 *[n]T。
*/
// 数组初始化
// 一维数组
// 全局
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str [5]string = [5]string{3: "hello world", 4: "tom"}

// 多维数组
// 全局
var brr0 [5][3]int
var brr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func main() {
	// 一维数组
	// 局部
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  int
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
	fmt.Println("-----------------------")
	// 多维数组
	// 局部
	a1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b1 := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	fmt.Println(brr0, brr1)
	fmt.Println(a1, b1)
	fmt.Println("-----------------------")
	// 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。
	a3 := [2]int{}
	fmt.Printf("a: %p\n", &a3)
	test(a3)
	fmt.Println(a3)
	fmt.Println("-----------------------")

	// 内置函数 len 和 cap 都返回数组长度 (元素数量)。
	a4 := [2]int{}
	println(len(a4), cap(a4))
	fmt.Println("-----------------------")

	// 多维数组遍历：
	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}
	fmt.Println("-----------------------")
	// 数组拷贝和传参
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
	fmt.Println("-----------------------")

	//数组练习
	//求数组所有元素之和
	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)
	//rand.Seed(time.Now().Unix())
	//var b2 [10]int
	//for i := 0; i < len(b2); i++ {
	//	// 产生一个0到1000随机数
	//	b[i] = rand.Intn(1000)
	//}
	//sum := sumArr(b2)
	//fmt.Printf("sum=%d\n", sum)
	fmt.Println("-----------------------")

	// 找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，
	// 找出两个元素之和等于8的下标分别是（0，4）和（1，2）
	b4 := [5]int{1, 3, 5, 8, 7}
	myTest(b4, 8)
	fmt.Println("-----------------------")

}

// 求元素和，是给定的值
func myTest(a [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

// 求元素和
func sumArr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

```