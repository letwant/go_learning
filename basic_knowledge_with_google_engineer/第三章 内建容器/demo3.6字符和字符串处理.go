package main

import "fmt"

func main (){
	s := "Yes, 你好，世界"  // UTF-8编码，可变长编码
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}

/*
使用len(str)会得到字节数，不会得到字符数，只有使用utf8.RuneCountInString才能获得字符数量
使用[]byte获得字节

其他字符串操作
Fields, Split, Join
Contains, Index
ToLower, ToUpper
Trim, TrimRight, TrimLeft
*/