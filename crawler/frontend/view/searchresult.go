package view

import (
	"html/template"
	"io"
	"learngo.com/crawler/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		// 用template.Must()包装一下，表示不允许Parse的时候出错
		template: template.Must(template.ParseFiles(filename)),
	}
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}