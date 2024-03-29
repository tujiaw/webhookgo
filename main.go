package main

import (
	"fmt"
	"net/http"
	"webhookgo/webhook"
)

func main() {
	webhook := webhook.New()
	webhook.AddCallback("tujiaw/webhookgo", func() {
		fmt.Println("tujiaw webhookgo")
	})
	webhook.AddCommand("tujiaw/ningtogo", "./scripts/ningtogo.sh")

	http.Handle("/webhook", webhook)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "success")
	})

	const ADDRESS = ":3333"
	fmt.Println("start address", ADDRESS)
	err := http.ListenAndServe(ADDRESS, nil)
	if err != nil {
		panic(err)
	}
}
