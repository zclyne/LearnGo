package parser

import (
	"learngo.com/crawler/engine"
	"regexp"
)

// 针对某一城市的Parser，获取城市下所有用户的信息

// 用户详细信息页面链接
var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
// 用户性别信息
var genderRe = regexp.MustCompile(`<td[^>]+><span[^>]+>性别：</span>([^>]+)</td>`)
// 下一页和其他相关链接
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	// 匹配用户名
	userMatches := profileRe.FindAllSubmatch(contents, -1)
	// 匹配用户性别
	genderMatches := genderRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for i, m := range userMatches { // 处理每一个match到的用户
		// 这里必须把要用到的url、name和gender都拷贝出来，因为ParseFunc并不是立即执行的，而当它执行时
		// for循环已经结束，对每一个ParseFunc调用，传入的m都是userMatches中的最后一项，导致内容出错
		// 而拷贝一份出来之后再传入，则可以避免这个问题
		url := string(m[1])
		name := string(m[2])
		gender := string(genderMatches[i][1])

		result.Requests = append(result.Requests, engine.Request{ // 用户页面url放入Requests中
			Url:        url,
			// 函数式编程方法，包装一下有多个参数的ParseProfile，实际调用ParseProfile时
			// 还要传入这里获得的用户名字和性别
			Parser: NewProfileParser(name, gender),
		})
	}

	// 匹配下一页链接和其他相关页面链接
	cityUrlMatches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range cityUrlMatches {
		cityUrl := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:        cityUrl,
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}