package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/hallo", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`hallo`))
	}))
	m.Handle("/rank", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`rank`))
	}))
	m.Handle("/history/xiaoqiang", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`history/xiaoqiang`))
	}))
	m.Handle("/history/xiaoding", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`history/xiaoding`))
	}))
	m.Handle("/history", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		qp := request.URL.Query()
		name := qp.Get("name")
		//if/else
		//switch/case
		//switch request.Method {
		//case "get"://...
		//case "post"://...
		//}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf(`%s,%s的健身房`, request.Method, name)))
	}))

	http.ListenAndServe(":8080", m)
}
