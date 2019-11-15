package main

import "fmt"

var (
	aa = 3
	bb = 4
	cc = 6
)
// 括号外不能使用:=

func main() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
	d, e, f := 1, 2, 5
	fmt.Println(d, e, f)
	fmt.Println("Hello World")
}

/*
变量定义总结：
	使用var关键字
		1、var a, b, c bool
		2、var s1, s2 string = "hello world"
		3、可放在函数内，或直接放在包内
		4、使用var()集中定义变量
	让编译器自动决定类型
		var a, b, i, s1, s2 = true, false, 3, "hello", "world"
	使用:=定义变量
		a, b, i, s1, s2 := true, false, 3, "hello", "world"
*/