package parser

import (
	"learngo.com/crawler/engine"
	"log"
	"regexp"
)

// 使用正则表达式匹配城市链接，[^>]*表示除了>以外，匹配任何字符
// 这里把>排除在外的原因是防止直接匹配到了</a>的>上
// 由于城市名字不为空，所以用[^<]+来匹配城市名
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 城市列表页的解析器
func ParseCityList(contents []byte) engine.ParseResult {
	log.Printf("Parsing City List...")
	re := regexp.MustCompile(cityListRe)

	// 执行匹配，返回值matches是[][]byte，其中每个[]byte都是一个城市对应的a标签字符串
	// matches := re.FindAll(contents, -1)

	// 由于要从匹配到的字符串里取出链接和城市名，所以要用FindAllSubmatch，返回值matches是[][][]byte
	// 其中最后一个[]byte是一个string，所以其效果和FindAllStringSubmatch返回的[][]stirng是相同的
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	//limit := 10 // 最多允许爬取的城市数
	for _, m := range matches { // 处理每一个match到的城市
		result.Requests = append(result.Requests, engine.Request{ // 城市url放入Requests中
			Url:        string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
		//limit--
		//if limit == 0 {
		//	break
		//}
	}
	// fmt.Println("Number of cities altogether:", len(matches)) // 470
	return result
}