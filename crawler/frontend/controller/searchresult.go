package controller

import (
	"context"
	"github.com/olivere/elastic/v7"
	"learngo.com/crawler/engine"
	"learngo.com/crawler/frontend/model"
	"learngo.com/crawler/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// 把数据从client中取出，然后送往view进行render
type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

// SearchResultHandler实现了net/http/handler interface
// localhost:8888/search?q=男 已购房&from=20
// from为分页的起点
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 获取参数
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	// 使用获取到的参数查询ElasticSearch，并渲染给view
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryStrng(q))).
		From(from).
		Do(context.Background())
	if err != nil { // 从ElasticSearch中搜索失败
		return result, err
	}

	// 搜索成功，填充result
	result.Hits = resp.TotalHits()
	result.Start = from
	for _, v := range resp.Each(reflect.TypeOf(engine.Item{})) {
		item := v.(engine.Item)
		result.Items = append(result.Items, item)
	}
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

// 在特定查询条件之前添加Payload.，例如原本查询条件是Age:(<30)，现在改成Payload.Age:(<30)
func rewriteQueryStrng(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	// $1就是括号内匹配到的部分
	return re.ReplaceAllString(q, "Payload.$1:")
}