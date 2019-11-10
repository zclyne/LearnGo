package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 限流器，防止爬虫爬取速度太快，触发目标网页的反爬虫机制
// 每100毫秒执行一次fetch
// 所有的worker都会使用同一个rateLimiter，worker总数由engine中的WorkerCount指定
var rateLimiter = time.Tick(10 * time.Millisecond)

// 根据url访问页面，并返回网页内容
func Fetch(url string) ([]byte, error) {
	//先等待rateLimiter中有信号接收到，然后再执行实际的爬取
	<- rateLimiter

	// 访问url
	// resp, err:= http.Get(url) // 直接get现在已经失效，因为网站做了反爬虫处理
	// 这里采用在header中加入User-Agent字段的方法来防止403 forbidden
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/71.0.3578.87 Safari/537.21")
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { // 访问出错,打印response的code
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 访问网页成功，返回resp的body中的内容
	return ioutil.ReadAll(resp.Body)
}