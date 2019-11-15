package main

import (
	"fmt"
	"my_go_demo/learn_go2/第三章 内建容器/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	q.Pop()
}
