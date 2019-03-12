package main

import "fmt"

func main() {
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	newSlice := slice[1:3]
	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为 60
	fmt.Println(slice, newSlice)

	newSlice = append(newSlice, 60) // 会修改底层共享数组对应位置的值，从而影响其他的slice
	fmt.Println(slice, newSlice)

	// slice[i:j:k]
	// 切片长度: j - i
	// 切片容量: k - i
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 对第三个元素做切片，并限制容量
	// 其长度和容量都是 1 个元素
	slice2 := source[2:3:3]
	// 向 slice 追加新字符串
	slice2 = append(slice2, "Kiwi")
	// 如果不加第三个索引，由于剩余的所有容量都属于slice2，向slice2追加Kiwi会改变原有底层数组索引为3的元素的值Banana。
	// 不过在上述中我们限制了slice2的容量为1。
	// 当我们第一次对slice2调用append的时候，会创建一个新的底层数组，这个数组包括2个元素，并将水果Plum 复制进来，
	// 再追加新水果Kiwi，并返回一个引用了这个底层数组的新切片， 从而不影响source的值。
	fmt.Println(source, slice2)

	// 切片在函数中传递：
	// 在64位架构的机器上，一个切片需要24字节的内存: 指针字段需要8字节，长度和容量字段分别需要8字节。
	// 由于与切片关联的数据包含在底层数组里，不属于切片本身，所以将切片复制到任意函数的时候，对底层数组大小都不会有影响。
	// 复制时只会复制切片本身，不会涉及底层数组
}
