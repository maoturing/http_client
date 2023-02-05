package main

import (
	"log"
	"net/http"
)

// 代码清单3-5  发送HEAD请求获取响应头
// $ curl --head http://localhost:18888
func main() {

	resp, err := http.Get("http://localhost:18888")

	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Header:", resp.Header)

	// head 请求的响应虽然只有head没有body, 但ContentLength的值仍为Get请求响应体的长度, 见https://stackoverflow.com/a/2773408/6528237
	log.Println("ContentLength:", resp.ContentLength)
}
