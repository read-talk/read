package handler

import (
	"fmt"
	readability "github.com/read-talk/readability/logic"
	"io/ioutil"
	"net/http"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("../../static/view/home.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("../../static/view/read.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	text := r.Form.Get("q")
	fmt.Println("接收到的参数text: ", text)
	// TODO: parse url
	content, err := readability.NewReadability(text, 30*time.Second)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	content.Parse()
	fmt.Println("解析参数得到: ", content.Title)
	w.Write([]byte(content.Title))
}
