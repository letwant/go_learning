package main

import (
	"fmt"
	"my_go_demo/learn_go2/第五章 面向接口/mock"
	"my_go_demo/learn_go2/第五章 面向接口/real"
)

/*
接口由使用者定义
*/
type Retriever interface {
	Get(url string) string
}


func download(r Retriever) string {
	return r.Get("http://www.immoc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake immoc.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}