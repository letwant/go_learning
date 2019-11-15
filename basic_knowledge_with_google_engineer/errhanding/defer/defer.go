package main

import (
	"bufio"
	"fmt"
	"my_go_demo/learn_go2/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(filename string) {
	file, err := os.OpenFile(
		filename, os.O_EXCL | os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()
	f := fib.Fibonacci()
	writer := bufio.NewWriter(file) // 这样较快一点
	defer writer.Flush()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
}
