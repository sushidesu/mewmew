package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type MewMewRequest struct {
	Message string `json:"message"`
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	fmt.Println("called: /send-message")

	if r.Method != "POST" {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
		return
	}

	var jsonBody MewMewRequest
	err := json.NewDecoder(r.Body).Decode(&jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send message to slack (#message-from-mewmew)
	webhook_url := os.Getenv("WEBHOOK_URL_MESSAGE_FROM_MEWMEW")

	// build message body
	type SlackSendMessage struct {
		Text string `json:"text"`
	}
	slack_message, err := json.Marshal(SlackSendMessage{
		Text: jsonBody.Message,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send request
	res, err := http.Post(webhook_url, "application/json", bytes.NewBuffer(slack_message))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	// response from slack
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(body)
}
