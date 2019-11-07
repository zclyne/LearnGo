package main

import (
	"fmt"
	"regexp"
)

const text = `My email is xxx@gmail.com
email1 is abc@def.org
email2 is      kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	// MustCompile表示传入的string必须符合正则表达式语法而不会有错误
	// 普通的Compile()方法有第二个参数err，表示传入的string是否符合正则表达式的语法
	// @之前不能用.+，因为.+会把空格也匹配进去
	// .匹配一个字符，+表示1个或多个；若用.*则表示0个或多个
	// 这里要用``，原因是普通的""包裹的字符串中，如果使用\.，则Go会认为这是一个转义字符，而不会只认为是一个.
	// 若要用""，则匹配一个.要写成\\.，而这种写法太丑，所以用``，表示没有任何的转义
	// [a-zA-Z0-9.]表示不仅允许匹配字母和数字，而且允许匹配.，因为有的邮箱在@之后有多个.，例如.com.cn
	// 在方括号内部不需要\.来转义，所以直接用.即可
	// 用()包裹的内容是正则表达式中要提取出来的部分，提取用re.FindAllSubmatch()或FindAllStringSubmatch()
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)

	// 寻找匹配的正则表达式，返回值match是一个string，表示查找到的匹配的字符串
	// match := re.FindString(text)

	// re.FindString()只返回匹配到的第一个，若要返回匹配到的所有结果，要用re.FindAllString
	// 第二个参数是要匹配的个数，传入-1表示匹配全部
	// 这里的match是一个[]string，存储所有匹配到的字符串
	// [xxx@gmail.com abc@def.org kkk@qq.com ddd@abc.com.cn]
	// match := re.FindAllString(text, -1)

	// 提取子匹配部分，返回值是一个二维的slice
	// 第0个是完整的匹配字符串，之后的是括号包裹起来的匹配到的内容
	// [[xxx@gmail.com xxx gmail com] [abc@def.org abc def org] [kkk@qq.com kkk qq com] [ddd@abc.com.cn ddd abc.com cn]]
	match := re.FindAllStringSubmatch(text, -1)

	for _, m := range match {
		fmt.Println(m)
	}

	fmt.Println(match)
}
