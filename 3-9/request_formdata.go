package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// 代码清单3-9  使用multipart/form-data形式发送文件, 参数会被分隔符boundary分隔开, 可以在echo_server中观察到
// 与Postman在Body->form-data中设置参数上传文件的作用相同, 请求中的参数会被分隔符boundary分割
// 与当前包内的request.html发送的请求相同, 表单中的参数都会被分隔符boundary分开并放在请求体中
// $ curl -F "name=Michael Jackson" -F "thumbnail=@photo.jpg" http://localhost:18888
func main() {
	var buffer bytes.Buffer
	// writer读取到的内容会保存到buffer中
	writer := multipart.NewWriter(&buffer)
	// 向输出流中写入字段
	writer.WriteField("name", "Michael Jackson")

	// 获取保存图片的输出流fileWriter
	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}

	// 读取图片到输入流readFile
	readFile, err := os.Open("D:\\MyProjects\\go\\src\\http_client\\3-9\\photo.jpg")
	if err != nil {
		panic(err) // 文件读取失败
	}
	defer readFile.Close()

	// 将readFile内容即图片拷贝到fileWriter中, 图片就保存到了输出流writer中
	io.Copy(fileWriter, readFile)
	writer.Close()

	// 发送POST请求, Content-Type为multipart/form-data, 表单内容包括字段图片都在buffer中
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	// 与上面的请求等价
	// http.Post("http://localhost:18888", "multipart/form-data; boundary=" + writer.Boundary(), &buffer)

	contentType := writer.FormDataContentType()
	log.Println("Content-Type:", contentType) // 打印这次请求的Content-Type

	if err != nil {
		panic(err) // 发送失败
	}
	log.Println("Status:", resp.Status) // 200 OK
}
