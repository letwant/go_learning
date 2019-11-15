package main

import "fmt"

func printSlice(s []int) {
	fmt.Println(s)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func main() {
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)

	s3 := make([]int, 10, 32) // 参数分别代表为元素类型，len长度，cap容积

	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) // 注意后面有三个点
	printSlice(s2) // 将8删除掉了

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(front, tail)
	printSlice(s2)
}

/*

*/