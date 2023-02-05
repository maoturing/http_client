package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// 代码清单3-3 GET请求的发送, 以及响应头, 响应体, 响应状态码的接收
// $ curl http://localhost:18888
func main() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	// resp.Body是一个输出流, 所以需要关闭
	defer resp.Body.Close()
	// ioutil.ReadAll()将响应流读取到一个字节数组中并返回
	// 因此body变量其实是byte[]类型
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	log.Println(string(body))									// 打印响应体 <html>...</html>
	log.Println("Status:", resp.Status)						// 200 OK
	log.Println("StatusCode:", resp.StatusCode)				// 200
	log.Println("ContentLength:", resp.ContentLength)		// 31 响应体长度, len("<html><body>hello</body></html>")的长度就是31

	log.Println("Headers:", resp.Header)
	log.Println("Content-Encoding:", resp.Header.Get("Content-Type"))
}
