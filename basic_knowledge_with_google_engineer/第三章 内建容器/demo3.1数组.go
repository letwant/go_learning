package main

import "fmt"

func printArray1(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100 // 无法改变数组的元素值
}

func printArray2(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100 // 无法改变数组的元素值
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10} // 让编译器自己定义数组长度

	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid) // [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]

	// 遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 使用range遍历数组，注意i是数组下标
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	printArray1(arr1)
	printArray1(arr3)
	fmt.Println(arr1, arr3) // [0 0 0 0 0] [2 4 6 8 10],值未修改

	printArray2(&arr1)
	printArray2(&arr3)
	fmt.Println(arr1, arr3) // [100 0 0 0 0] [100 4 6 8 10], 值已修改
}

/*
数组是值类型，数组作为参数传给函数时，会拷贝一份，其他大部分语言作为数组传参会传一个引用，能改变数组的值
go语言一般不直接使用数组，而使用切片
*/