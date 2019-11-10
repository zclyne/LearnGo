package persist

import (
	"context"
	"encoding/json"
	"learngo.com/crawler/model"
	"testing"
	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
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
	}

	// 把profile插入到ElasticSearch中
	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	// 使用id从ElasticSearch中获取插入的profile
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	// 反序列化
	var actual model.Profile
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v, created %v", actual, expected)
	}

}