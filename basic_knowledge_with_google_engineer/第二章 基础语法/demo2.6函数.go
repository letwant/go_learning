package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator:" + op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)", opName, a, b)
	return op(a, b)
}

func sumArgs(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func main() {
	if result, err := eval(3, 4, "/"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(
			float64(a), float64(b)))
	}, 3, 4))
	fmt.Println(sumArgs(1, 2, 3, 4, 5))
}


/*
go语言只有可变参数列表
go语言没有函数重载，没有默认参数，没有可选参数
*/

