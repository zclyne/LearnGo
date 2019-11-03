package main

import "fmt"

func main() {
	// map[K]V，此处key的类型和value的类型都是string
	// map可以复合，例如map[K1]map[K2]V，表示key的类型为K1，value的类型是一个map，该map的key的类型和value的类型为K2和V
	m := map[string]string {
		"name": "yifan",
		"gender": "male",
		"age": "21",
	}
	fmt.Println(m) // map[age:21 gender:male name:yifan]

	// 另一种创建map的方式，可以定义空map
	m2 := make(map[string]int) // m2 == empty map
	fmt.Println(m2) // map[]

	// 第三种方式，注意和第二种方式的区别
	// go语言的nil和其他语言的null不同，nil可以参与运算
	var m3 map[string]int // m3 == nil
	fmt.Println(m3) // map[]

	// 遍历map中的元素，同样可以用_来替代k或v
	// 遍历map时，各个pair出现的顺序并不固定，这是因为这个map是一个hashmap，内部元素无序
	fmt.Println("Traversing map...")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	name := m["name"]
	fmt.Println(name) // yifan
	nume := m["nume"]
	fmt.Println(nume) // 空字符串，因为空字符串是string的zero value，而在go中，如果从map中取不到对应元素，就返回zero value
	fmt.Println(m2["asd"]) // 0，原因同上
	// 判断key是否在map中
	age, ok1 := m["age"]
	nama, ok2 := m["nama"]
	fmt.Println(age, ok1) // 21 true
	fmt.Println(nama, ok2) // 空字符串 false
	// 常用写法
	if gendre, ok3 := m["ok3"]; ok3 {
		fmt.Println(gendre)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"] // 虽然name已经定义过，但是ok没定义过，所以仍可以用:=
	fmt.Println(name, ok)
	delete(m, "name") // yifan true
	name, ok = m["name"]
	fmt.Println(name, ok) // 空字符串 false

	// 用作key的条件
	// map使用哈希表，因此用作key的类型必须可以比较是否相等
	// 除了slice，map，function以外的内建类型都可以作为key
	// 自定义的struct若不包含上述字段，则也可以作为key
}
