package main

import "fmt"

func main() {
	// 声明第一个包含 5 个元素的字符串数组 var array1 [5]string
	var array1 [5]string
	// 声明第二个包含 5 个元素的字符串数组
	// 用颜色初始化数组
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	// 把 array2 的值复制到 array1
	array1 = array2
	fmt.Println(array1, array2)
	array2[0] = "red"
	fmt.Println(array1, array2)

	var a1 [3]*string
	a2 := [3]*string{new(string), new(string), new(string)}
	*a2[0] = "a"
	*a2[1] = "b"
	*a2[2] = "c"

	a1 = a2
	fmt.Println(*a1[0], *a2[0])
	*a2[0] = "A"
	fmt.Println(*a1[0], *a2[0])
}
