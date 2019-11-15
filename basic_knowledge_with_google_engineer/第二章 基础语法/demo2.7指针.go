package main

import "fmt"

func swap(a, b *int) {
	*b, *a = *a, *b
}


func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
	d, e := 3, 4
	swap(&d, &e)
	fmt.Println(d, e)
}

/*
指针不能运算
值传递，引用传递？
go语言只有值传递这一种方式
*/