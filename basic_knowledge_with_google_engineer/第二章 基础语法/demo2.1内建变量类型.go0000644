package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// rune 字符型，go语言的char类型，32位的
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1) // (0+1.2246467991473515e-16i)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1) // (0.000+0.000i) ,只要小数点后三位
}

func main () {
	euler()
	a, b := 3, 4
	var c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

/*
go语言没有隐式类型转换，只有强制类型转换
var c int = math.Sqrt(a*a + b*b) go语言报错，其他语言有隐式转换会自动将float类型的数据转换为int型数据
*/