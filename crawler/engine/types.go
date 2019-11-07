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
	Items []interface{} // Items可以是任何类型
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}