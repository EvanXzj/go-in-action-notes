package main

import "fmt"

func main() {
	// 通过声明映射创建一个 nil 映射
	// var colors map[string]string
	// 将 Red 的代码加入到映射
	// colors["Red"] = "#da1337"
	// Error:
	// Runtime Error:
	// panic: runtime error: assignment to entry in nil map

	// 初始化map
	cols := make(map[string]string)
	cols["Red"] = "#da1337"
	cols["Coral"] = "#ff7F50"
	cols["DarkGray"] = "#a9a9a9"
	// 检测key是否存在
	value, exists := cols["Red"]
	if exists {
		fmt.Println("Exists")
		fmt.Println("Value: ", value)
	}
	// 删除key
	delete(cols, "Coral")
	value, exists = cols["Coral"]
	if !exists {
		fmt.Println("Coral has been deleted")
	}

	// range
	for key, value := range cols {
		fmt.Printf("Key: %s \\ Value: %s\n", key, value)
	}

	fmt.Println()
	// 在函数间传递
	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
	// 调用函数来移除指定的键
	removeColor(colors, "Coral")
	fmt.Println()
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

// removeColor将指定映射里的键删除
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
