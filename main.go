package main

import "net/http"

type MyHandler struct{}

func (m MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("received"))
}

func main() {
	handler := MyHandler{}
	http.HandleFunc("/", handler.ServeHTTP)
	http.ListenAndServe("0.0.0.0:3000", nil)
}
