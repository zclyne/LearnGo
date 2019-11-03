package mock

import "fmt"

// 注意这里的Retriever虽然实现了main.go中的Retriever接口，但是没有出现这个关键字
// 在main.go中，只要传入的Retriever有Get()方法，就认为是合法的

type Retriever struct {
	Contents string
}

// 实现Stringer接口
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}

// Retriever同时实现了Poster接口和Retriever接口

// 因为这个方法中对Contents内容做了修改，所以要用指针接收者
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
