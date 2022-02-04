package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", myhandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func myhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	fmt.Println("请求方法：", r.Method)
	fmt.Println("url: ", r.URL)
	fmt.Println("请求头：", r.Header)
	fmt.Println("请求消息体：", r.Body)

	w.Write([]byte("收到啦"))
}
