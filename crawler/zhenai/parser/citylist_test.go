package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// 这里不应该采用访问网络来获取要测试的数据的方式，因为测试机可能无法联网，或网络故障
	// contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	// 这里采用把网页html内容保存到本地文件，然后从文件上读取的方式来测试Parser
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s\n", contents)

	result := ParseCityList(contents)
	const resultSize = 470
	expectedUrls := []string {
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	// 验证数量是否正确
	if len(result.Requests) != resultSize {
		t.Errorf("result.Requests should have %d requests, but had %d", resultSize, len(result.Requests))
	}

	// 验证城市url和城市名是否正确
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expted url #%d: %s, but was %s", i, url, result.Requests[i].Url)
		}
	}

}