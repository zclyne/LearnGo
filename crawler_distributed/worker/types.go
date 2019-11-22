package worker

import (
	"learngo.com/crawler/engine"
	"learngo.com/crawler/zhenai/parser"
	"learngo.com/crawler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

// 由于engine.Request和engine.ParseResult中都含有interface{}，无法在网络上流式传输
// 所以在这里建立了rpc中单独的Request和ParseResult来包装它们
type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Items []engine.Item
	Requests []Request
}

// 把engine.Request和engine.ParseResult转换成rpc的Request和ParseResult
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url:    r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult {
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

// 反序列化
func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}
// 负责把rpc的SerializedParser转换回engine.Parser
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		// p.Args是一个[]string类型，其中第一个是username，第二个是usergender
		}
		return parser.NewProfileParser(p.Args.([]string)[0], p.Args.([]string)[1]), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items:    r.Items,
	}
	for _, req := range r.Requests {
		engineRequest, err := DeserializeRequest(req)
		if err != nil { // 由于request数量非常多，所以出现一个error无关紧要，打一句log记录一下并跳过这个request即可
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineRequest)
	}
	return result
}