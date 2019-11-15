package main

import "fmt"

func updateSlice(s []int)  {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6]) // [2, 3, 4, 5]
	fmt.Println("arr[:6] =", arr[:6])   // [0, 1, 2, 3, 4, 5]
	fmt.Println("arr[2:] =", arr[2:])   // [2, 3, 4, 5, 6, 7]
	fmt.Println("arr[:] =", arr[:])     // [0, 1, 2, 3, 4, 5, 6, 7]

	s1 := arr[2:]
	s2 := arr[:]

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr) // [0 1 100 3 4 5 6 7],值修改，说明切片Slice是引用传递

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	arr[0], arr[2] = 0, 2
	s3 := arr[2:6]
	s4 := s3[3:5]
	fmt.Println(s3)  // [2, 3, 4, 5]
	fmt.Println(s4)

	fmt.Println("arr  =", arr)
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))

	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr1[2:6]
	s2 = s1[3:5]
	s3 = append(s2, 10)
	// s3和s4不是原来的arr
	s4 = append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s1, s2, s3, s4, s5, arr)
}

/*
数组转切片: var arr [5]int  -->  slice := arr[:]  后面加一个[:]即可
slice可以向后扩展，不可以向前扩展
s[i]不可以超越len(s)，向后扩展不可以超越递增数组cap(s)
向Slice添加元素
添加元素时如果超越cap，系统会重新分配更大的底层数组
由于值传递的关系，必须接收append的返回值
s = append(s, val)
*/