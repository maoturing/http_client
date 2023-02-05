package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)

	log.Println("start http listening: 18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	// 打印请求request
	fmt.Println(string(dump))

	// 将字符串写入到输出流中
	fmt.Fprint(w, "<html><body>hello</body></html>")
}
