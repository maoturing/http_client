package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// 代码清单3-9
// $ curl -F "name=Michael Jackson" -F "thumbnail=@photo.jpg" http://localhost:18888
func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err) // 文件读取失败
	}
	defer readFile.Close()

	io.Copy(fileWriter, readFile)
	writer.Close()

	contentType := writer.FormDataContentType()
	log.Println("Content-Type:", contentType)		// 打印这次请求的Content-Type

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err) // 发送失败
	}
	log.Println("Status:", resp.Status) // 200 OK
}
