package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WebhookHandler struct {
	monitor map[string]func()
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{
		monitor: make(map[string]func()),
	}
}

func (webhook *WebhookHandler) AddHandler(fullname string, handler func()) {
	webhook.monitor[fullname] = handler
}

func (webhook *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	repository := payload["repository"].(map[string]interface{})
	fullname := repository["full_name"].(string)
	if webhook.monitor[fullname] != nil {
		webhook.monitor[fullname]()
		fmt.Fprint(w, "success")
	} else {
		fmt.Fprint(w, "not handler")
	}
}

func main() {
	webhook := NewWebhookHandler()
	webhook.AddHandler("tujiaw/webhookgo", func() {
		fmt.Println("1111111111111111111")
	})
	webhook.AddHandler("tujiaw/ningtogo", func() {
		fmt.Println("22222222222222222")
	})

	http.Handle("/webhook", webhook)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "success")
	})

	const ADDRESS = ":3000"
	fmt.Println("start address", ADDRESS)
	err := http.ListenAndServe(ADDRESS, nil)
	if err != nil {
		panic(err)
	}
}
