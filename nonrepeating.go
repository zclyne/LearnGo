package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	// 以字符为key，最后出现的位置为value，建立map
	lastOccurred := make(map[rune]int)
	start, maxLength := 0, 0
	// 注意这里要对[]rune(s)进行遍历，而不能直接range s
	// 直接range s时，ch正常，但是i在遇到中文时会每次递增3
	for i, ch := range []rune(s) {
		// 注意这里要判断lastI是否存在
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start { // 当前字符上一次出现的位置在start之后，要更新start
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("bbtablud"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcbacbbab"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	// 以下两个计算结果不正确，因为中文字符不止8bit，不能把string强制转换成byte数组之后操作
	fmt.Println(lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
}
