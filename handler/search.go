package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
		r.ParseForm()
		text := r.Form.Get("q")
		fmt.Println("接收到的参数text: ", text)
		// TODO: parse url
		w.Write([]byte(text))
	}
}
