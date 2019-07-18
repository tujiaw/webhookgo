package main

import (
	"fmt"
	"net/http"
)

type WebhookHandler struct {
	monitorList []string
}

func (handler *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ref := r.PostFormValue("ref")
	fmt.Println("ref", ref)
}

func main() {
	const ADDRESS = ":3000"

	http.Handle("/webhook", &WebhookHandler{})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "success")
	})

	fmt.Println("start address", ADDRESS)
	err := http.ListenAndServe(ADDRESS, nil)
	if err != nil {
		panic(err)
	}
}
