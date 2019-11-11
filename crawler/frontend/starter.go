package main

import (
	"learngo.com/crawler/frontend/controller"
	"net/http"
)

func main() {
	// 处理根目录，如果没有这一步，则页面将无法获取css和js文件
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))
	// 处理/search路径，用于前端展示结果
	http.Handle("/search", controller.CreateSearchResultHandler("crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}