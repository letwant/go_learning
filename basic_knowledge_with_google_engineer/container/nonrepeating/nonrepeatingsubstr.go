package main

import "fmt"

/*
abcabcbb --> abc
bbbbbb   --> b
pwwkew   --> wke
*/

/*
对于每一个字母x

lastOccurred[x]不存在，或者<start -->无需操作
lastOccurred[x] >= start --> 更新start
更新Occurred[x]，更新maxLength
*/
var lastOccurred = make([]int, 0xffff)
func lengthOfNonRepeatingSubStr(s string) int {
	// stores last occurred pos + 1.
	// 0 means not seen
	for i := range lastOccurred {
		lastOccurred[i] = 0
	}
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI > start {
			start = lastI
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwkefa"))
	fmt.Println(lengthOfNonRepeatingSubStr("你好世界"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三三二一"))

}