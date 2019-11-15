package main

import "fmt"

/*
结构体只支持封装，不支持继承和多态
只有struct，没有class
*/

type point struct {
	i, j int
}

type treeNode struct {
	value int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value:value} // 局部变量地址外面也能使用，相较之，c++中就不可以使用局部变量的地址
}

func (node treeNode) printNode()  {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) { // 直接使用指针就可以修改值了
	if node == nil {
		fmt.Println("setting value to nil node. Ignored")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.printNode()
	node.right.traverse()
}

func main() {
	var root treeNode
	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) // go语言中没有->，只有'.'，因为用了语法糖
	root.left.right = createNode(2)
	root.right.left.setValue(4)

	root.traverse()


	var pRoot *treeNode
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.printNode()
}

/*
不论地址还是结构本身，一律使用.来访问成员
注意返回了局部变量的地址
不需要知道内存分配在堆上还是栈上
只有使用指针才可以改变结构内容
nil指针也可以调用方法

值接收者VS指针接收者
要改变内容必须使用指针接收者
结构过大也考虑使用指针接收者 ，传指针消耗内存小
一致性：如有指针接收者，最好都是指针接收者

值接收者是go语言特有的
值/指针接收者均可接收值/指针
*/


