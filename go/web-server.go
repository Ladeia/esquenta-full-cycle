package main

import "net/http"

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8888", nil)
}

func HelloWorld(W http.ResponseWriter, r *http.Request) {
	W.Write([]byte("Hello World!"))
}
