package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 根据url访问页面，并返回网页内容
func Fetch(url string) ([]byte, error) {
	// 访问url
	resp, err:= http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { // 访问出错,打印response的code
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 访问网页成功，返回resp的body中的内容
	return ioutil.ReadAll(resp.Body)
}