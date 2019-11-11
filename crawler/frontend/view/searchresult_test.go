package view

import (
	"learngo.com/crawler/engine"
	"learngo.com/crawler/frontend/model"
	common "learngo.com/crawler/model" // 因为同时有2个model存在，所以这里要取个别名，common
	"os"
	"testing"
)

// 模板测试

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Id: "1866830740",
		Url: "https://album.zhenai.com/u/1866830740",
		Type: "zhenai",
		// 用别名common来调用包中的内容
		Payload: common.Profile{
			Name:          "花儿少年",
			Gender:        "男士",
			Age:           23,
			Height:        175,
			Weight:        67,
			Income:        "5-8千",
			Marriage:      "未婚",
			Education:     "大专",
			AncestralHome: "重庆",
			Constellation: "射手座",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}