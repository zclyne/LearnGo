package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// http client

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://cn.bing.com", nil)
	// 使用header来模拟手机的访问
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	// 直接使用http访问网页
	// resp, err := http.Get("http://cn.bing.com")

	// 使用默认的client访问网页
	// resp, err := http.DefaultClient.Do(request)

	// 自定义client，打印重定向过程
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
