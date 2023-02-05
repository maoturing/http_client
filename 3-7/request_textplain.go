package main

import (
	"log"
	"net/http"
	"os"
)

// 代码清单3-7
// $ curl -F "name=Michael Jackson" -F "thumbnail=@photo.jpg" http://localhost:18888
func main() {

	file, err := os.Open("D:\\MyProjects\\go\\src\\http_client\\resource\\a.txt")
	if err != nil {
		panic(err) // 文件读取失败
	}

	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err) // 发送失败
	}
	log.Println("Status:", resp.Status) // 200 OK
}
