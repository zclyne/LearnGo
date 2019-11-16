package persist

import (
	"context"
	"encoding/json"
	"learngo.com/crawler/engine"
	"learngo.com/crawler/model"
	"testing"
	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Id: "1866830740",
		Url: "https://album.zhenai.com/u/1866830740",
		Type: "zhenai",
		Payload: model.Profile{
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

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// 把profile插入到ElasticSearch中
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// 使用id从ElasticSearch中获取插入的profile
	client, err = elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	// 反序列化
	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	// 由于Item的Payload是interface{}，所以直接反序列化会得到一个map，而不是Profile
	// 所以定义了方法FromJsonObj来把Payload转换成Profile，然后再放回actual.Payload
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v, created %v", actual, expected)
	}

}