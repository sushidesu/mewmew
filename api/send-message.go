package handler

import (
	"encoding/json"
	"net/http"
)

type sendMessageResponse struct {
	Message string                 `json:"message"`
	Body    map[string]interface{} `json:"body"`
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
		return
	}

	var jsonBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(sendMessageResponse{
		Message: "hello",
		Body:    jsonBody,
	})
}
