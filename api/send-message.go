package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type sendMessageResponse struct {
	Message string `json:"message"`
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if r.Method != "POST" {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
		return
	}

	var jsonBody map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send message to slack (#message-from-mewmew)
	webhook_url := os.Getenv("WEBHOOK_URL_MESSAGE_FROM_MEWMEW")

	// build message body
	body, err := json.Marshal(jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	type SlackSendMessage struct {
		Text string `json:"text"`
	}
	slack_message := SlackSendMessage{
		Text: string(body),
	}
	message, err := json.Marshal(slack_message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send request
	res, err := http.NewRequest("POST", webhook_url, bytes.NewBuffer(message))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	json.NewEncoder(w).Encode(sendMessageResponse{
		Message: "ok",
	})
}
