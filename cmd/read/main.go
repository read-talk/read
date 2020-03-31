package main

import (
	"fmt"
	"net/http"

	"github.com/read-talk/read/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("../../static"))))
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/search", handler.SearchHandler)
	fmt.Printf("服务启动中，开始监听端口 [%s]... \n", ":2020")
	http.ListenAndServe(":2020", mux)
}
