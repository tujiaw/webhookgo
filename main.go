package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WebhookHandler struct {
	monitorList []string
}

func (handler *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	var payload map[string]interface{}
	err = json.Unmarshal(bbody, &payload)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Println(payload)
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
