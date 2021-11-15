package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/png"
	"net/http"

	"github.com/sushidesu/mewmew/lib/circle"
)

type DetectionRequest struct {
	Image string `json:"image"`
}

func DetectionHandler(w http.ResponseWriter, r *http.Request) {
	// allow CORS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "invalid Content-Type", http.StatusBadRequest)
		return
	}

	// parse json
	var jsonBody DetectionRequest
	err := json.NewDecoder(r.Body).Decode(&jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate
	if jsonBody.Image == "" {
		http.Error(w, "image is required", http.StatusBadRequest)
		return
	}

	// get dots
	b, err := base64.StdEncoding.DecodeString(jsonBody.Image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dots := circle.CreateDots(circle.GetPtsFromImage(img))

	// response
	bounds := img.Bounds()
	result := circle.ShowCircle(*dots, bounds)
	var buf bytes.Buffer
	err = png.Encode(&buf, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	str := base64.StdEncoding.EncodeToString(buf.Bytes())
	w.Write([]byte("data:image/png;base64," + str))
}
