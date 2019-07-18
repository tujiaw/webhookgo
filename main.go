package main

import (
	"fmt"
	"net/http"
)

type WebhookHandler struct {
	monitorList []string
}

func (handler *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "success")
}

func main() {
	const ADDRESS = ":3000"

	http.Handle("/", &WebhookHandler{})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "world")
	})

	fmt.Println("start address", ADDRESS)
	err := http.ListenAndServe(ADDRESS, nil)
	if err != nil {
		panic(err)
	}
}
