package handler

import (
	"encoding/json"
	"net/http"
)

type sendMessageResponse struct {
	Message string `json:"message"`
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(sendMessageResponse{
		Message: "hello",
	})
}
