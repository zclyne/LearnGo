package engine

// 统一定义所使用的类型

// request结构体包含要访问的url，以及用来解析该url内容的parser
type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

// ParseResult中包含所有该页面下的待访问的链接，以及相应的项目名
// 例如对于城市列表页，Requests中包含所有城市的url
// Items中包含所有的城市名
type ParseResult struct {
	Requests []Request
	Items []Item // Items中的元素必须是Item类型
}

// id、type和url是各种不同爬虫所爬取的内容都应该具有的公共部分，所以把它抽象出来放入type中
// type是特定爬虫所爬取的内容类型信息，此处是"zhenai"，表示爬的网站
// Payload则是因爬虫不同而异的数据部分，例如此处就是用户的Profile
type Item struct {
	Url string
	Id string
	Type string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}