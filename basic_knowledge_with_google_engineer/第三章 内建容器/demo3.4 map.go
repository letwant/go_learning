package main

import "fmt"

func main() {
	// 创建map的集中方法
	m := map[string]string {
		"name": "ccmouse",
		"course": "golang",
		"site": "immoc",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int) // m2 == empty map
	fmt.Println(m2)

	var m3 map[string]int // m3 == nil
	fmt.Println(m3)

	fmt.Println("Traversing map") // map是无序的
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	// 如果key不存在与map中，则取出的值为空, zeroValue
	if causeName, ok := m["cause"]; ok { // 较为标准的取值方法
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
