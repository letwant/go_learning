package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scaner := bufio.NewScanner(file)

	for scaner.Scan() {
		fmt.Println(scaner.Text()) // 读一行
	}
}

func forever () {
	for {  // 相当于while
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		)
}

/*
for, if 后面的条件没有括号
if条件里也可以定义变量
switch里面可以没有break
switch后可以没有表达式
没有while
*/