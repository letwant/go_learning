package main

import "fmt"

/*
const filename = "abc.txt"
const 数值可作为各种类型使用
const a, b = 3, 4
*/

func enums() {
	const (
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
	const (
		c = iota
		javascript
		_
		django
	)
	fmt.Println(c, javascript, django)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb) // 1 1024 1048576 1073741824 1099511627776 1125899906842624
}

func main() {
	enums()
}

/*变量定义要点回顾
1、没有char，只有rune
2、原生支持复数类型
*/