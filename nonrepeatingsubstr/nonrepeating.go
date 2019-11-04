package nonrepeatingsubstr

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	// 以字符为key，最后出现的位置为value，建立map
	// lastOccurred := make(map[rune]int)
	// 用map速度比较慢，所以这里用空间换时间，开一个很大的数组
	// 创建数组的过程可以放到函数外部，进一步优化性能
	lastOccurred := make([]int, 0xffff)
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	start, maxLength := 0, 0

	// 使用map时的写法
	// 注意这里要对[]rune(s)进行遍历，而不能直接range s
	// 直接range s时，ch正常，但是i在遇到中文时会每次递增3
	//for i, ch := range []rune(s) {
	//	// 注意这里要判断lastI是否存在
	//	if lastI, ok := lastOccurred[ch]; ok && lastI >= start { // 当前字符上一次出现的位置在start之后，要更新start
	//		start = lastI + 1
	//	}
	//	if i - start + 1 > maxLength {
	//		maxLength = i - start + 1
	//	}
	//	lastOccurred[ch] = i
	//}

	// 使用数组时的写法
	for i, ch := range []rune(s) {
		// 把字符作为索引传入，会自动把字符转换成对应的UTF-8编码对应的整数，因此作为索引是合法的
		// lastI == -1表示还没有出现过这个字符
		if lastI:= lastOccurred[ch]; lastI != -1 && lastI >= start { // 当前字符上一次出现的位置在start之后，要更新start
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
