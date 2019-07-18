package webhook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

type WebhookHandler struct {
	callback map[string]func()
	command  map[string]string
}

func New() *WebhookHandler {
	return &WebhookHandler{
		callback: make(map[string]func()),
		command:  make(map[string]string),
	}
}

func (webhook *WebhookHandler) AddCallback(fullname string, handler func()) {
	webhook.callback[fullname] = handler
}

func (webhook *WebhookHandler) AddCommand(fullname string, command string) {
	webhook.command[fullname] = command
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
	success := false
	cb := webhook.callback[fullname]
	if cb != nil {
		go cb()
		success = true
	}

	cmd := webhook.command[fullname]
	if len(cmd) > 0 {
		go exec.Command("/bin/bash", "-c", cmd)
		success = true
	}

	if success {
		fmt.Fprint(w, "success")
	} else {
		fmt.Fprint(w, "not handler")
	}
}
