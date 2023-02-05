package main

import (
	"log"
	"net/http"
	"net/url"
)

// 代码清单3-6 使用POST请求发送x-www-form-urlencoded形式的表单
// 与Postman在Body->x-www-form-urlencoded中设置参数的作用相同, 参数都会编码后放在请求体中
// 与当前包内的request.html发送的请求相同, 表单中的参数都会编码后放在请求体中
// $ curl -d test=value http://localhost:18888
func main() {
	values := url.Values{
		"test": {"value"},
	}
	// 发送POST请求, 以form表单的形式传递参数, 参数不会在url中体现, 参数在请求体中, 默认contentType为x-www-form-urlencoded
	resp, err := http.PostForm("http://localhost:18888", values)
	// 与上面这行代码等价, 设置Content-Type为application/x-www-form-urlencoded, 对参数进行编码
	// resp, err := http.Post("http://localhost:18888", "application/x-www-form-urlencoded", strings.NewReader(values.Encode())

	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status) // 200 OK
}
