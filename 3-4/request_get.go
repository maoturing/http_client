package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// 代码清单3-4 使用Get方法发送查询
// 与Postman中设置Params-QueryParams作用相同, 参数都会拼接到url后面
// $ curl -G --data-urlencode "query=hello world" http://localhost:18888
func main() {
	values := url.Values{
		"query": {"hello wolrd"},
	}
	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
