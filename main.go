package main

import (
	"fmt"
	"github_project/web/model"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// http服务
	http.HandleFunc("/index", model.Index)
	http.HandleFunc("/gen", model.Tpl)

	// 使用9000端口
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		fmt.Println("HTTP server start failed , err %v", err)
		return
	}
}
