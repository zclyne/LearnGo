package main

import (
	"LearnGo/retriever/mock"
	"LearnGo/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	// 在interface内部，函数名之前不需要加func关键字
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string {
		"name": "yifan",
		"course": "golang",
	})
}

// 接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	// 这里既可以调用s.Get()，也可以调用s.Post()以及s.Connect()
	s.Post(url, map[string]string {
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting")
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	// 生成一个mock.Retriever，然后拷贝给r
	r = &mock.Retriever{"this is a fake imooc.com"} // 此时r的类型是mock.Retriever
	inspect(r) // mock.Retriever {this is a fake imooc.com}

	// real.Retriever的Get()方法接收者是*Retriever
	// 生成一个real.Retriever，然后把指向这个Retriever的指针给r
	r = &real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute} // 此时r的类型是*real.Retriever
	inspect(r) // *real.Retriever {Mozilla/5.0 1m0s}

	// Type assertion，通过. + (类型)来获得其interface内部的实际类型
	// realRetriever := r.(real.Retriever) // 报错，impossible type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut) // 1m0s

	// mockRetriever := r.(mock.Retriever) // 报错，main.Retriever is *real.Retriever, not mock.Retriever
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("Not a mock retriever")
	}

	// fmt.Println(download(r))
	// 接口变量也可以看过一个struct，其中包括实现者的类型和实现者的值，或实现者的类型和实现者的指针
	// 由于接口本身可以包含实现者的指针，所以一般不使用接口自身的指针
	// 指针接收者实现智能以指针方式使用；值接收者两者都可

	fmt.Println("Try a session")
	s := &mock.Retriever{Contents: "This is another mock.retriever"}
	// 因为mock.Retriever同时实现了Poster和Retriever接口，所以可以传入session方法中
	fmt.Println(session(s)) // another faked imooc.com

	// s实现了Stringer接口，可以使用.String()方法
	fmt.Println("Test String()")
	fmt.Println(s.String())
}
