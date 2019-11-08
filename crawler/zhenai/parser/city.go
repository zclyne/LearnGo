package parser

import (
	"LearnGo/crawler/engine"
	"regexp"
)

// 针对某一城市的Parser，获取城市下所有用户的信息

// 提取用户页面链接的正则表达式
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const genderRe = `<td[^>]+><span[^>]+>性别：</span>([^>]+)</td>`

func ParseCity(contents []byte) engine.ParseResult {
	// 匹配用户名
	re := regexp.MustCompile(cityRe)
	userMatches := re.FindAllSubmatch(contents, -1)
	// 匹配用户性别
	re = regexp.MustCompile(genderRe)
	genderMatches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for i, m := range userMatches { // 处理每一个match到的用户
		// 这里必须把要用到的url、name和gender都拷贝出来，因为ParseFunc并不是立即执行的，而当它执行时
		// for循环已经结束，对每一个ParseFunc调用，传入的m都是userMatches中的最后一项，导致内容出错
		// 而拷贝一份出来之后再传入，则可以避免这个问题
		url := string(m[1])
		name := string(m[2])
		gender := string(genderMatches[i][1])

		result.Items = append(result.Items, "User " + name) // 用户名字名字放在Items中
		result.Requests = append(result.Requests, engine.Request{ // 用户页面url放入Requests中
			Url:        url,
			// 函数式编程方法，包装一下有多个参数的ParseProfile，实际调用ParseProfile时
			// 还要传入这里获得的用户名字和性别
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, gender)
			},
		})
	}
	return result
}