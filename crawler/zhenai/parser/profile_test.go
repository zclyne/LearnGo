package parser

import (
	"learngo.com/crawler/engine"
	"learngo.com/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "https://album.zhenai.com/u/1866830740", "花儿少年", "男士")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; but got %v", result.Items)
	}

	// profile := result.Items[0].(model.Profile)
	actual := result.Items[0]

	expected := engine.Item{
		Id: "1866830740",
		Type: "zhenai",
		Url: "https://album.zhenai.com/u/1866830740",
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

	if actual != expected {
		t.Errorf("expected %v; but got %v", expected, actual)
	}

}